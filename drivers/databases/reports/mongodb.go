package reports

import (
	"context"
	"disspace/business/reports"
	"disspace/helpers/messages"

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
