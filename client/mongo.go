package client

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type MongoDep struct {
	Client          *mongo.Client
	ExampleDatabase *mongo.Database
}

func NewMongoClient(ctx context.Context) (*MongoDep, error) {
	opts := options.
		Client().
		ApplyURI("mongodb://localhost:27017")

	ctxWithTimeout, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctxWithTimeout, opts)
	if err != nil {
		return nil, err
	}

	go func() {
		<-ctx.Done()
		err = client.Disconnect(context.Background())
		if err != nil {
			log.Fatalf("error disconecting mongodb client. Error: %s", err.Error())
		}
	}()

	log.Printf("connected to mongo database")

	return &MongoDep{
		Client:          client,
		ExampleDatabase: client.Database("example-app-db"),
	}, nil
}
