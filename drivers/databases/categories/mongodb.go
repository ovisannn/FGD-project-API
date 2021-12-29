package categories

import (
	"context"
	"disspace/business/categories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// fmt.Println(result)
	return result, nil
}

func (repository *MongoDBCategoriesRepository) Create(ctx context.Context, data *categories.Domain) (categories.Domain, error) {
	category := FromDomain(*data)

	cursor, err := repository.Conn.Collection("categories").InsertOne(ctx, category)
	if err != nil {
		return categories.Domain{}, err
	}

	categoryID := cursor.InsertedID.(primitive.ObjectID).Hex()
	return categories.Domain{ID: categoryID}, nil
}

func (repository *MongoDBCategoriesRepository) GetByID(ctx context.Context, id string) (categories.Domain, error) {
	result := Categories{}
	convert, errConvert := primitive.ObjectIDFromHex(id)
	if errConvert != nil {
		return categories.Domain{}, errConvert
	}
	find := bson.D{{Key: "_id", Value: convert}}
	err := repository.Conn.Collection("categories").FindOne(ctx, find).Decode(result)

	if err != nil {
		panic(err)
	}

	return result.ToDomain(), nil
}
