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
	// Start Error Handling
	_, errConvId := primitive.ObjectIDFromHex(id)
	if errConvId != nil {
		return messages.ErrInvalidUserID
	}

	// End Error Handling

	report := FromDomain(*reportDomain)

	_, err := repository.Conn.Collection("reports").InsertOne(ctx, report)
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
