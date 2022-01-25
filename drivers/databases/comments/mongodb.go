package comments

import (
	"context"
	"disspace/business/comments"
	"disspace/helpers/messages"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBCommentRepository struct {
	Conn *mongo.Database
}

func NewMongoDBCommentRepository(conn *mongo.Database) comments.Repository {
	return &MongoDBCommentRepository{
		Conn: conn,
	}
}

var userLookup = bson.M{
	"from":         "user_profile",
	"localField":   "username",
	"foreignField": "username",
	"as":           "user",
}

var votesLookup = bson.M{
	"from":         "votes",
	"localField":   "comment_id",
	"foreignField": "reference_id",
	"as":           "votes",
}

var commentLookup = bson.M{
	"from":         "comments",
	"localField":   "comment_id",
	"foreignField": "parent_id",
	"as":           "comments",
}

var convIdToString = bson.M{"comment_id": bson.M{"$toString": "$_id"}}
var countVotes = bson.M{"num_votes": bson.M{"$sum": "$votes.status"}}
var countComments = bson.M{"num_comments": bson.M{"$size": "$comments"}}

func (repository *MongoDBCommentRepository) Create(ctx context.Context, commentDomain *comments.Domain, id string) (comments.Domain, error) {
	// Start Error Handling
	// _, errConvId := primitive.ObjectIDFromHex(id)
	// if errConvId != nil {
	// 	return comments.Domain{}, messages.ErrInvalidUserID
	// }

	// // Filter
	// filter := bson.D{{Key: "_id", Value: convId}}

	// // Check if user and reference exist in database
	// checkUser, errCheckUser := repository.Conn.Collection("users").CountDocuments(ctx, filter)
	// if errCheckUser != nil {
	// 	return errCheckUser
	// }
	// if checkUser == 0 {
	// 	return comments.Domain{}, messages.ErrUnauthorizedUser
	// }

	// End Error Handling

	comment := FromDomain(*commentDomain)

	cursor, err := repository.Conn.Collection("comments").InsertOne(ctx, comment)
	if err != nil {
		return comments.Domain{}, err
	}

	commentId := cursor.InsertedID.(primitive.ObjectID).Hex()
	return comments.Domain{ID: commentId}, nil
}

func (repository *MongoDBCommentRepository) Delete(ctx context.Context, id string, threadId string) error {
	// Start Error Handling
	_, errConvId := primitive.ObjectIDFromHex(id)
	if errConvId != nil {
		return messages.ErrInvalidUserID
	}

	_, errConvThreadId := primitive.ObjectIDFromHex(threadId)
	if errConvThreadId != nil {
		return messages.ErrInvalidThreadID
	}
	// End Error Handling

	filter := bson.D{{Key: "$and", Value: []interface{}{
		bson.D{{Key: "user_id", Value: id}}, bson.D{{Key: "thread_id", Value: threadId}},
	}}}
	result, err := repository.Conn.Collection("comments").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return messages.ErrDataNotFound
	}
	return nil
}

func (repository *MongoDBCommentRepository) Search(ctx context.Context, q string, sort string) ([]comments.Domain, error) {
	var result []comments.Domain
	if sort == "" {
		sort = "created_at"
	}
	sorting := bson.M{sort: -1}
	opts := options.Find().SetSort(sorting)

	// Create index for comments collection
	model := mongo.IndexModel{Keys: bson.D{{Key: "text", Value: "text"}}}
	_, errIndex := repository.Conn.Collection("comments").Indexes().CreateOne(ctx, model)
	if errIndex != nil {
		return []comments.Domain{}, errIndex
	}

	// Search
	filter := bson.D{}
	if q != "" {
		filter = bson.D{{Key: "username", Value: primitive.Regex{Pattern: q, Options: "i"}}}
	}

	cursor, err := repository.Conn.Collection("comments").Find(ctx, filter, opts)
	if err != nil {
		return []comments.Domain{}, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return []comments.Domain{}, err
	}
	return result, nil
}

func (repository *MongoDBCommentRepository) GetByID(ctx context.Context, id string) (comments.Domain, error) {
	result := []Comment{}

	conv, errConvId := primitive.ObjectIDFromHex(id)
	if errConvId != nil {
		return comments.Domain{}, messages.ErrInvalidUserID
	}

	filter := bson.D{{Key: "_id", Value: conv}}

	query := []bson.M{
		{
			"$match": filter,
		},
		{
			"$addFields": convIdToString,
		},
		{
			"$lookup": userLookup,
		},
		{
			"$lookup": votesLookup,
		},
		{
			"$lookup": commentLookup,
		},
		{
			"$unwind": "$user",
		},
		{
			"$addFields": countVotes,
		},
		{
			"$addFields": countComments,
		},
	}

	cursor, err := repository.Conn.Collection("comments").Aggregate(ctx, query)
	if err != nil {
		return comments.Domain{}, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return comments.Domain{}, err
	}

	var res = result[0]

	return res.ToDomain(), nil
}

func (repository *MongoDBCommentRepository) GetAllInThread(ctx context.Context, threadId string, parentId string) ([]comments.Domain, error) {
	result := []comments.Domain{}

	// filter := bson.D{{Key: "$and", Value: []interface{}{
	// 	bson.D{{Key: "thread_id", Value: threadId}}, bson.D{{Key: "parent_id", Value: parentId}},
	// }}}
	filter := bson.D{{Key: "thread_id", Value: threadId}}

	query := []bson.M{
		{
			"$match": filter,
		},
		{
			"$addFields": convIdToString,
		},
		{
			"$lookup": userLookup,
		},
		{
			"$lookup": votesLookup,
		},
		{
			"$lookup": commentLookup,
		},
		{
			"$unwind": "$user",
		},
		{
			"$addFields": countVotes,
		},
		{
			"$addFields": countComments,
		},
	}

	cursor, err := repository.Conn.Collection("comments").Aggregate(ctx, query)
	if err != nil {
		return []comments.Domain{}, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return []comments.Domain{}, err
	}

	return result, nil
}
