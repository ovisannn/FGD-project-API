package mongoDB

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(ctx context.Context) (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://admin:26QulskduMYF9ns9@disspace-shard-00-00.vltti.mongodb.net:27017/myFirstDatabase?ssl=true&replicaSet=atlas-e260ru-shard-0&authSource=admin&retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	db := client.Database("disspace")
	return db, nil
}
