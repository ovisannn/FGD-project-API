package moderators

import (
	"context"
	"disspace/business/moderators"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBModeratorsRepository struct {
	Conn *mongo.Database
}

func NewMongoDBModeratorsRepository(conn *mongo.Database) moderators.Repository {
	return &MongoDBModeratorsRepository{
		Conn: conn,
	}
}
func (repository *MongoDBModeratorsRepository) GetByCategoryID(ctx context.Context, idCategory string) ([]moderators.Domain, error) {
	result := []moderators.Domain{}
	filter := bson.D{{Key: "categori_id", Value: idCategory}}
	cursor, err := repository.Conn.Collection("moderators").Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &result); err != nil {
		return []moderators.Domain{}, err
	}
	// fmt.Println(result)
	// fmt.Println(idCategory)
	return result, nil

}
