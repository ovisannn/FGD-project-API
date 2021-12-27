package threads

import (
	"context"
	"disspace/business/threads"

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

	return result, nil
}

func (repository *MongoDBThreadRepository) Create(ctx context.Context, threadDomain *threads.Domain) (threads.Domain, error) {
	thread := FromDomain(*threadDomain)

	cursor, err := repository.Conn.Collection("threads").InsertOne(ctx, thread)
	if err != nil {
		return threads.Domain{}, err
	}

	threadId := cursor.InsertedID.(primitive.ObjectID).Hex()
	return threads.Domain{ID: threadId}, nil
}

func (repository *MongoDBThreadRepository) GetByID(ctx context.Context, id string) (threads.Domain, error) {
	result := Thread{}

	convert, errorConvert := primitive.ObjectIDFromHex(id)
	if errorConvert != nil {
		return threads.Domain{}, errorConvert
	}

	filter := bson.D{{Key: "_id", Value: convert}}
	err := repository.Conn.Collection("threads").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result.ToDomain(), nil
}



























func (repository *MongoDBThreadRepository) Update(ctx context.Context, threadDomain *threads.Domain, id string) error {
	thread := FromDomain(*threadDomain)

	convert, errorConvert := primitive.ObjectIDFromHex(id)
	if errorConvert != nil {
		return errorConvert
	}

	update := bson.D{{Key: "$set", Value: thread}}
	_, err := repository.Conn.Collection("threads").UpdateByID(ctx, convert, update)
	if err != nil {
		return err
	}
	return nil
}
