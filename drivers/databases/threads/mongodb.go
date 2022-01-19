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

var userLookup = bson.M{
	"from":         "user_profile",
	"localField":   "user_id",
	"foreignField": "username",
	"as":           "user",
}

func (repository *MongoDBThreadRepository) GetAll(ctx context.Context, sort string) ([]threads.Domain, error) {
	var result []threads.Domain

	countVotes := bson.M{"num_votes": bson.M{"$sum": "$votes.status"}}
	countComments := bson.M{"num_comments": bson.M{"$size": "$comments"}}
	convIdToString := bson.M{"thread_id": bson.M{"$toString": "$_id"}}
	if sort == "" {
		sort = "created_at"
	}
	sorting := bson.M{sort: -1}

	query := []bson.M{
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
		{
			"$sort": sorting,
		},
	}

	cursor, err := repository.Conn.Collection("threads").Aggregate(ctx, query)
	if err != nil {
		return []threads.Domain{}, err
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
	thread := FromDomainUpdate(*threadDomain)

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

func (repository *MongoDBThreadRepository) Search(ctx context.Context, q string, sort string) ([]threads.Domain, error) {
	var result []threads.Domain

	// Create index for threads collection
	model := mongo.IndexModel{Keys: bson.D{{Key: "title", Value: "text"}, {Key: "content", Value: "text"}}}
	_, errIndex := repository.Conn.Collection("threads").Indexes().CreateOne(ctx, model)
	if errIndex != nil {
		return []threads.Domain{}, errIndex
	}

	// Search
	countVotes := bson.M{"num_votes": bson.M{"$sum": "$votes.status"}}
	countComments := bson.M{"num_comments": bson.M{"$size": "$comments"}}
	convIdToString := bson.M{"thread_id": bson.M{"$toString": "$_id"}}
	if sort == "" {
		sort = "created_at"
	}
	sorting := bson.M{sort: -1}
	title := bson.M{"title": primitive.Regex{Pattern: q, Options: "i"}}

	query := []bson.M{
		{
			"$match": title,
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
			"$unwind": "user",
		},
		{
			"$addFields": countVotes,
		},
		{
			"$addFields": countComments,
		},
		{
			"$sort": sorting,
		},
	}

	cursor, err := repository.Conn.Collection("threads").Aggregate(ctx, query)
	if err != nil {
		return []threads.Domain{}, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return []threads.Domain{}, err
	}
	return result, nil
}
