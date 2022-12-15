package repository

import (
	"context"
	"github.com/henriquegmendes/go-base-api-rest/models"
)

type ExampleRepository interface {
	Create(ctx context.Context, example *models.Example) error
}
