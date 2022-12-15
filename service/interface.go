package service

import (
	"context"
	"github.com/henriquegmendes/go-base-api-rest/dtos/request"
	"github.com/henriquegmendes/go-base-api-rest/dtos/response"
)

type ExampleService interface {
	Create(ctx context.Context, exampleRequest request.ExampleRequest) (*response.ExampleResponse, error)
}
