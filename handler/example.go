package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	internalErrors "github.com/henriquegmendes/go-base-api-rest/errors"
	"github.com/henriquegmendes/go-base-api-rest/server/helpers/router"
	"io"
	"net/http"
)

func LoadExampleRoutes(internalRouter router.InternalRouter) {
	internalRouter.POST("/example", Create)
}

type Request struct {
	Name string `json:"name"`
}

func Create(ctx *gin.Context) (*router.InternalResponse, error) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		return nil, err
	}

	var requestBody Request
	err = json.Unmarshal(bodyBytes, &requestBody)
	if err != nil {
		return nil, internalErrors.NewApplicationError("invalid json body request", http.StatusBadRequest)
	}

	// call service...

	return router.NewInternalResponse(requestBody, http.StatusCreated), nil
}
