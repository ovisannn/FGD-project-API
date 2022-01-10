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

var votesLookup = bson.M{
	"from":         "votes",
	"localField":   "thread_id",
	"foreignField": "reference_id",
	"as":           "votes",
}

var commentLookup = bson.M{
	"from":         "comments",
	"localField":   "thread_id",
	"foreignField": "thread_id",
	"as":           "comments",
}

func (repository *MongoDBThreadRepository) GetAll(ctx context.Context) ([]threads.Domain, error) {
	var result []threads.Domain

	countVotes := bson.M{"num_votes": bson.M{"$sum": "$votes.status"}}
	countComments := bson.M{"num_comments": bson.M{"$size": "$comments"}}
	convIdToString := bson.M{"thread_id": bson.M{"$toString": "$_id"}}

	query := []bson.M{
		{
			"$addFields": convIdToString,
		},
		{
			"$lookup": votesLookup,
		},
		{
			"$lookup": commentLookup,
		},
		{
			"$addFields": countVotes,
		},
		{
			"$addFields": countComments,
		},
	}

	cursor, err := repository.Conn.Collection("threads").Aggregate(ctx, query)
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

func (repository *MongoDBThreadRepository) Delete(ctx context.Context, id string) error {
	convert, errorConvert := primitive.ObjectIDFromHex(id)
	if errorConvert != nil {
		return errorConvert
	}
	filter := bson.D{{Key: "_id", Value: convert}}
	_, err := repository.Conn.Collection("threads").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (repository *MongoDBThreadRepository) GetByID(ctx context.Context, id string) (threads.Domain, error) {
	result := []Thread{}

	convert, errorConvert := primitive.ObjectIDFromHex(id)
	if errorConvert != nil {
		return threads.Domain{}, errorConvert
	}

	filter := bson.D{{Key: "_id", Value: convert}}



	countVotes := bson.M{"num_votes": bson.M{"$sum": "$votes.status"}}
	countComments := bson.M{"num_comments": bson.M{"$size": "$comments"}}
	convIdToString := bson.M{"thread_id": bson.M{"$toString": "$_id"}}

	query := []bson.M{
		{
			"$match": filter,
		},
		{
			"$addFields": convIdToString,
		},
		{
			"$lookup": votesLookup,
		},
		{
			"$lookup": commentLookup,
		},
		{
			"$addFields": countVotes,
		},
		{
			"$addFields": countComments,
		},
	}

	cursor, err := repository.Conn.Collection("threads").Aggregate(ctx, query)
	if err != nil {
		return threads.Domain{}, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return threads.Domain{}, err
	}
	
	var res = result[0]

	return res.ToDomain(), nil
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