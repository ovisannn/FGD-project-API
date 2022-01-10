package comments

import (
	"context"
	"disspace/business/comments"
	"disspace/helpers/messages"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
	_, errConvId := primitive.ObjectIDFromHex(id)
	if errConvId != nil {
		return comments.Domain{}, messages.ErrInvalidUserID
	}

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
