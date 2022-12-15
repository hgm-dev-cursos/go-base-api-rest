package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
			// error handling mais decente
			ctx.JSON(500, gin.H{
				"message": fmt.Sprintf("error in handler: %s", err.Error()),
			})
			return
		}

		// success response handler
		ctx.JSON(response.StatusCode, response.Body)
	}
}
