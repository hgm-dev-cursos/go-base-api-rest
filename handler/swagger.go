package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
)

func LoadSwaggerRoutes(ginRouter *gin.Engine) {
	swaggerUIOpts := middleware.SwaggerUIOpts{SpecURL: "/swagger.json"}
	swaggerHandler := middleware.SwaggerUI(swaggerUIOpts, nil)

	ginRouter.GET("/docs", func(ctx *gin.Context) {
		swaggerHandler.ServeHTTP(ctx.Writer, ctx.Request)
	})
	ginRouter.StaticFile("/swagger.json", "./swagger.json")
}
