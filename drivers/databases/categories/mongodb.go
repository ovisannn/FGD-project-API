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
	convert, errorConvert := primitive.ObjectIDFromHex(id)
	if errorConvert != nil {
		return categories.Domain{}, errorConvert
	}
	filter := bson.D{{Key: "_id", Value: convert}}
	err := repository.Conn.Collection("categories").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result.ToDomain(), nil
}

func (repository *MongoDBCategoriesRepository) Delete(ctx context.Context, id string) error {
	convert, errConvert := primitive.ObjectIDFromHex(id)
	if errConvert != nil {
		return errConvert
	}
	_, err := repository.Conn.Collection("categories").DeleteOne(ctx, bson.D{{Key: "_id", Value: convert}})

	if err != nil {
		return err
	}

	return nil
}

func (repository *MongoDBCategoriesRepository) Update(ctx context.Context, data *categories.Domain, id string) error {
	category := FromDomain(*data)

	convert, errConv := primitive.ObjectIDFromHex(id)
	if errConv != nil {
		return errConv
	}
	_, err := repository.Conn.Collection("categories").UpdateByID(ctx, convert, bson.D{{Key: "$set", Value: category}})
	if err != nil {
		return err
	}
	return nil
}
