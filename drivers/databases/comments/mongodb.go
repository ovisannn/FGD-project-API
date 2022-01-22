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
	result := Comment{}

	conv, errConvId := primitive.ObjectIDFromHex(id)
	if errConvId != nil {
		return comments.Domain{}, messages.ErrInvalidUserID
	}
	var filter = bson.D{{Key: "_id", Value: conv}}
	err := repository.Conn.Collection("comments").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return comments.Domain{}, err
	}

	return result.ToDomain(), nil
}
