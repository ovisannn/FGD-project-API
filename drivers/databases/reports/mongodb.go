package reports

import (
	"context"
	"disspace/business/reports"
	"disspace/helpers/messages"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBReportRepository struct {
	Conn *mongo.Database
}

func NewMongoDBReportRepository(conn *mongo.Database) reports.Repository {
	return &MongoDBReportRepository{
		Conn: conn,
	}
}

func (repository *MongoDBReportRepository) Create(ctx context.Context, reportDomain *reports.Domain, id string) error {
	report := FromDomain(*reportDomain)
	col := "reports"

	if report.TargetType == 1 {
		col = "user_reports"
	}

	_, err := repository.Conn.Collection(col).InsertOne(ctx, report)
	if err != nil {
		return messages.ErrDuplicatedData
	}
	return nil
}

func (repository *MongoDBReportRepository) GetAll(ctx context.Context, sort string) ([]reports.Domain, error) {
	var result []reports.Domain
	if sort == "" {
		sort = "count"
	}
	sorting := bson.M{sort: -1}

	var threadLookup = bson.M{
		"from":         "threads",
		"localField":   "thread_id",
		"foreignField": "_id",
		"as":           "thread",
	}

	convThread := bson.M{"thread_id": bson.M{"$toObjectId": "$target_id"}}
	// convThread := bson.M{
	// 	"input": "$target_id",
	// 	"to": "objectId",
	// }

	query := []bson.M{
		{
			"$group": bson.M{
				"_id":         bson.M{"filter1": "$target_id", "filter2": "$target_type", "filter3": "$description"},
				"target_id":   bson.M{"$first": "$target_id"},
				"target_type": bson.M{"$first": "$target_type"},
				"description": bson.M{"$first": "$description"},
				"count":       bson.M{"$sum": 1},
			},
		},
		{
			"$addFields": convThread,
		},
		{
			"$lookup": threadLookup,
		},
		{
			"$unwind": "$thread",
		},
		{
			"$sort": sorting,
		},
	}

	cursor, err := repository.Conn.Collection("reports").Aggregate(ctx, query)
	if err != nil {
		return []reports.Domain{}, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return []reports.Domain{}, err
	}

	return result, nil
}

func (repository *MongoDBReportRepository) GetUserReport(ctx context.Context, sort string, q string) ([]reports.Domain, error) {
	var result []reports.Domain
	if sort == "" {
		sort = "count"
	}
	sorting := bson.M{sort: -1}
	username := bson.M{"target_id": primitive.Regex{Pattern: q, Options: "i"}}

	query := []bson.M{
		{
			"$match": username,
		},
		{
			"$group": bson.M{
				"_id":         bson.M{"filter1": "$target_id", "filter3": "$description"},
				"target_id":   bson.M{"$first": "$target_id"},
				"target_type": bson.M{"$first": "$target_type"},
				"description": bson.M{"$first": "$description"},
				"count":       bson.M{"$sum": 1},
			},
		},
		{
			"$sort": sorting,
		},
	}

	cursor, err := repository.Conn.Collection("user_reports").Aggregate(ctx, query)
	if err != nil {
		return []reports.Domain{}, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return []reports.Domain{}, err
	}

	return result, nil
}

func (repository *MongoDBReportRepository) GetCommentReport(ctx context.Context, sort string, q string) ([]reports.Domain, error) {
	var result []reports.Domain
	if sort == "" {
		sort = "count"
	}
	sorting := bson.M{sort: -1}

	var commentLookup = bson.M{
		"from":         "comments",
		"localField":   "comment_id",
		"foreignField": "_id",
		"as":           "comment",
	}

	var userLookup = bson.M{
		"from":         "user_profile",
		"localField":   "comment.username",
		"foreignField": "username",
		"as":           "comment.user",
	}

	convComment := bson.M{"comment_id": bson.M{"$toObjectId": "$target_id"}}

	query := []bson.M{
		{
			"$group": bson.M{
				"_id":         bson.M{"filter1": "$target_id", "filter2": "$target_type", "filter3": "$description"},
				"target_id":   bson.M{"$first": "$target_id"},
				"target_type": bson.M{"$first": "$target_type"},
				"description": bson.M{"$first": "$description"},
				"count":       bson.M{"$sum": 1},
			},
		},
		{
			"$addFields": convComment,
		},
		{
			"$lookup": commentLookup,
		},
		{
			"$unwind": "$comment",
		},
		{
			"$lookup": userLookup,
		},
		{
			"$unwind": "$comment.user",
		},
		{
			"$sort": sorting,
		},
	}

	cursor, err := repository.Conn.Collection("reports").Aggregate(ctx, query)
	if err != nil {
		return []reports.Domain{}, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return []reports.Domain{}, err
	}

	return result, nil
}
