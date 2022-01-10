package user

import (
	"context"
	"disspace/business/user"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBUserRepository struct {
	Conn *mongo.Database
}

func NewMongoDBUserRepository(conn *mongo.Database) user.Repository {
	return &MongoDBUserRepository{
		Conn: conn,
	}
}

func (repository *MongoDBUserRepository) Register(ctx context.Context, data *user.UserDomain) (user.UserDomain, error) {
	// gs://disspace-250a1.appspot.com/profile_pict/profile_default.jpg
	newUser := UserFromDomain(*data)
	cursor, err := repository.Conn.Collection("users").InsertOne(ctx, newUser)
	userID := cursor.InsertedID.(primitive.ObjectID).Hex()
	newProfileUser := UserProfile{
		UserId:      userID,
		ProfilePict: "gs://disspace-250a1.appspot.com/profile_pict/profile_default.jpg",
		Bio:         " ",
		Following:   []string{},
		Followers:   []string{},
		Threads:     []string{},
		Reputation:  0,
	}
	_, errProfile := repository.Conn.Collection("user_profile").InsertOne(ctx, newProfileUser)
	if err != nil {
		return user.UserDomain{}, err
	}
	if errProfile != nil {
		return user.UserDomain{}, err
	}
	return user.UserDomain{}, nil

}

func (repository *MongoDBUserRepository) UserProfileGetByUserID(ctx context.Context, id string) (user.UserProfileDomain, error) {
	result := UserProfile{}
	filter := bson.D{{Key: "user_id", Value: id}}
	err := repository.Conn.Collection("user_profile").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		panic(err)
	}
	return result.UserProfileToDomain(), nil
}
