package votes

import (
	"context"
	"disspace/business/votes"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBVoteRepository struct {
	Conn *mongo.Database
}

func NewMongoDBVoteRepository(conn *mongo.Database) votes.Repository {
	return &MongoDBVoteRepository{
		Conn: conn,
	}
}

func (repository *MongoDBVoteRepository) Store(ctx context.Context, voteDomain *votes.Domain, id string) error {
	vote := FromDomain(*voteDomain)

	_, err := repository.Conn.Collection("likes").InsertOne(ctx, vote)
	if err != nil {
		return err
	}
	return nil
}
