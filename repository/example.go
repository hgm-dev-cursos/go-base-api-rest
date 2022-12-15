package repository

import (
	"context"
	"github.com/henriquegmendes/go-base-api-rest/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type exampleRepository struct {
	Database       *mongo.Database
	CollectionName string
}

func NewExampleRepository(database *mongo.Database) ExampleRepository {
	return &exampleRepository{
		Database:       database,
		CollectionName: "example",
	}
}

func (r *exampleRepository) Create(ctx context.Context, example *models.Example) error {
	_, err := r.Database.Collection(r.CollectionName).InsertOne(ctx, example)
	return err
}
