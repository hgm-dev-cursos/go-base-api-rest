package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/henriquegmendes/go-base-api-rest/dependencies"
	"github.com/henriquegmendes/go-base-api-rest/dtos/request"
	"github.com/henriquegmendes/go-base-api-rest/helpers"
	"github.com/henriquegmendes/go-base-api-rest/server/helpers/router"
	"github.com/henriquegmendes/go-base-api-rest/service"
	"io"
	"net/http"
)

type exampleHandler struct {
	exampleService service.ExampleService
}

func LoadExampleRoutes(globalDeps *dependencies.GlobalDeps, internalRouter router.InternalRouter) {
	handler := exampleHandler{
		exampleService: globalDeps.ExampleService,
	}

	internalRouter.POST("/example", handler.Create)
}

func (h *exampleHandler) Create(ctx *gin.Context) (*router.InternalResponse, error) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	var requestBody request.ExampleRequest
	err = helpers.UnmarshalAndValidate(bodyBytes, &requestBody)
	if err != nil {
		return nil, err
	}

	response, err := h.exampleService.Create(ctx, requestBody)
	if err != nil {
		return nil, err
	}

	return router.NewInternalResponse(response, http.StatusCreated), nil
}
