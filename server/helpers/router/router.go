package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	internalErrors "github.com/henriquegmendes/go-base-api-rest/errors"
	"net/http"
)

type internalRouter struct {
	GinRouter *gin.Engine
	BasePath  string
}

func NewInternalRouter(ginRouter *gin.Engine, basePath string) InternalRouter {
	return &internalRouter{
		GinRouter: ginRouter,
		BasePath:  basePath,
	}
}

func (r *internalRouter) POST(path string, handler InternalHandler) {
	r.GinRouter.POST(r.BasePath+path, r.handle(handler))
}

func (r *internalRouter) handle(internalHandler InternalHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := internalHandler(ctx)
		if err != nil {
			r.handleApplicationError(ctx, err)
			return
		}

		r.handleSuccessResponse(ctx, response)

	}
}

func (r *internalRouter) handleApplicationError(ctx *gin.Context, err error) {
	var applicationError *internalErrors.ApplicationError
	if errors.As(err, &applicationError) {
		ctx.JSON(applicationError.StatusCode, applicationError)
		return
	}

	ctx.JSON(http.StatusInternalServerError, internalErrors.DefaultApplicationError)
}

func (r *internalRouter) handleSuccessResponse(ctx *gin.Context, response *InternalResponse) {
	if response == nil {
		ctx.JSON(http.StatusNoContent, nil)
		return
	}

	ctx.JSON(response.StatusCode, response.Body)
}
