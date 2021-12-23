package threads

import (
	"context"
	"disspace/business/threads"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBThreadRepository struct {
	Conn *mongo.Database
}

func NewMongoDBThreadRepository(conn *mongo.Database) threads.Repository {
	return &MongoDBThreadRepository{
		Conn: conn,
	}
}

func (repository *MongoDBThreadRepository) GetAll(ctx context.Context) ([]threads.Domain, error) {
	var result []threads.Domain
	cursor, err := repository.Conn.Collection("threads").Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &result); err != nil {
		return []threads.Domain{}, err
	}

	fmt.Println(result)
	return result, nil
}

func (repository *MongoDBThreadRepository) GetByID(ctx context.Context, id string) (threads.Domain, error) {
	result := Thread{}

	convert, errorConvert := primitive.ObjectIDFromHex(id)
	if errorConvert != nil {
		return threads.Domain{}, errorConvert
	}

	err := repository.Conn.Collection("threads").FindOne(ctx, bson.D{{Key: "_id", Value: convert}}).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result.ToDomain(), nil
}
