package service

import (
	"context"
	"github.com/henriquegmendes/go-base-api-rest/dtos/request"
	"github.com/henriquegmendes/go-base-api-rest/dtos/response"
	internalErrors "github.com/henriquegmendes/go-base-api-rest/errors"
	"github.com/henriquegmendes/go-base-api-rest/models"
	"github.com/henriquegmendes/go-base-api-rest/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type exampleService struct {
	exampleRepository repository.ExampleRepository
}

func NewExampleService(exampleRepository repository.ExampleRepository) ExampleService {
	return &exampleService{
		exampleRepository: exampleRepository,
	}
}

func (s *exampleService) Create(ctx context.Context, exampleRequest request.ExampleRequest) (*response.ExampleResponse, error) {
	nowUTC := time.Now().UTC()
	newExample := &models.Example{
		ID:        primitive.NewObjectID(),
		Name:      exampleRequest.Name,
		CreatedAt: nowUTC,
		UpdatedAt: nowUTC,
	}

	err := s.exampleRepository.Create(ctx, newExample)
	if err != nil {
		return nil, internalErrors.NewApplicationError("error saving new example", http.StatusInternalServerError)
	}

	return &response.ExampleResponse{
		ID:   newExample.ID.Hex(),
		Name: newExample.Name,
	}, nil
}
