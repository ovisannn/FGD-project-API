package threads

import (
	"context"
	"disspace/business/threads"

	"go.mongodb.org/mongo-driver/bson"
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
	
	return result, nil
}
