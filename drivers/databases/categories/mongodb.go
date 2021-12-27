package categories

import (
	"context"
	"disspace/business/categories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBCategoriesRepository struct {
	Conn *mongo.Database
}

func NewMongoDBCategoriesRepository(conn *mongo.Database) categories.Repository {
	return &MongoDBCategoriesRepository{
		Conn: conn,
	}
}

func (repository *MongoDBCategoriesRepository) GetAll(ctx context.Context) ([]categories.Domain, error) {
	var result []categories.Domain
	cursor, err := repository.Conn.Collection("categories").Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &result); err != nil {
		return []categories.Domain{}, err
	}

	return result, nil
}
