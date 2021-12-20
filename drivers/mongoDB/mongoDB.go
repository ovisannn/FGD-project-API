package mongoDB

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect_to_db(ctx context.Context) (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb+srv://admin:26QulskduMYF9ns9@disspace.vltti.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("disspace"), nil
}
