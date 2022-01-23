package votes

import (
	"context"
	"disspace/business/votes"
	"disspace/helpers/messages"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// Start Error Handling

	// Check User ID and Reference ID validity
	// _, errConvId := primitive.ObjectIDFromHex(id)
	// if errConvId != nil {
	// 	return messages.ErrInvalidUserID
	// }

	_, errConvRefId := primitive.ObjectIDFromHex(voteDomain.ReferenceID)
	if errConvRefId != nil {
		return messages.ErrInvalidReferenceID
	}

	// // Filter
	// filter1 := bson.D{{Key: "_id", Value: convId}}
	// filter2 := bson.D{{Key: "_id", Value: convRefId}}

	// // Check if user and reference exist in database
	// checkUser, errCheckUser := repository.Conn.Collection("users").CountDocuments(ctx, filter1)
	// if errCheckUser != nil {
	// 	return errCheckUser
	// }
	// if checkUser == 0 {
	// 	return messages.ErrUnauthorizedUser
	// }

	// checkThread, errCheckThread := repository.Conn.Collection("threads").CountDocuments(ctx, filter2)
	// if errCheckThread != nil {
	// 	return errCheckThread
	// }

	// checkCom, errCheckCom := repository.Conn.Collection("comments").CountDocuments(ctx, filter2)
	// if errCheckCom != nil {
	// 	return errCheckCom
	// }
	// if checkThread == 0 && checkCom == 0 {
	// 	return messages.ErrReferenceNotFound
	// }

	// End Error Handling

	vote := FromDomain(*voteDomain)

	_, err := repository.Conn.Collection("votes").InsertOne(ctx, vote)
	if err != nil {
		return err
	}
	return nil
}

func (repository *MongoDBVoteRepository) Update(ctx context.Context, status int, id string, refid string) error {
	// Start Error Handling

	// Check User ID and Reference ID validity
	// _, errConvId := primitive.ObjectIDFromHex(id)
	// if errConvId != nil {
	// 	return messages.ErrInvalidUserID
	// }

	_, errConvRefId := primitive.ObjectIDFromHex(refid)
	if errConvRefId != nil {
		return messages.ErrInvalidReferenceID
	}

	// End Error Handling
	opts := options.Update().SetUpsert(true)

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: status}}}}
	filter := bson.D{{Key: "$and", Value: []interface{}{
		bson.D{{Key: "username", Value: id}}, bson.D{{Key: "reference_id", Value: refid}},
	}}}

	result, err := repository.Conn.Collection("votes").UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	// Check if data exist in the database
	if result.MatchedCount == 0 {
		return messages.ErrDataNotFound
	}
	return nil
}

func (repository *MongoDBVoteRepository) GetIsVoted(ctx context.Context, username string, refId string) (votes.Domain, error) {
	result := Vote{}

	var filter = bson.D{{Key: "$and", Value: []interface{}{
		bson.D{{Key: "username", Value: username}},
		bson.D{{Key: "reference_id", Value: refId}},
	}}}
	err := repository.Conn.Collection("votes").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return votes.Domain{}, err
	}

	return result.ToDomain(), nil
}
