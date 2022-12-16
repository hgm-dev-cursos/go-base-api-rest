package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	internalErrors "github.com/henriquegmendes/go-base-api-rest/errors"
	"github.com/henriquegmendes/go-base-api-rest/server/helpers/router"
	"net/http"
)

const authorizationHeader = "Authorization"

func AuthRouteMiddleware(internalHandler router.InternalHandler) router.InternalHandler {
	return func(ctx *gin.Context) (*router.InternalResponse, error) {
		err := validateHeader(ctx, authorizationHeader)
		if err != nil {
			return nil, err
		}

		return internalHandler(ctx)
	}
}

func AuthMiddleware(ctx *gin.Context) error {
	return validateHeader(ctx, authorizationHeader)
}

func validateHeader(ctx *gin.Context, headerName string) error {
	authHeader := ctx.GetHeader(headerName)
	if authHeader == "" {
		return internalErrors.NewApplicationError(fmt.Sprintf("missing %s header", headerName), http.StatusBadRequest)
	}
	if authHeader != "super-secret-auth" {
		return internalErrors.NewApplicationError(fmt.Sprintf("invalid %s header", headerName), http.StatusUnauthorized)
	}

	return nil
}
