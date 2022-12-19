package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/henriquegmendes/go-base-api-rest/dependencies"
	"github.com/henriquegmendes/go-base-api-rest/handler"
	"github.com/henriquegmendes/go-base-api-rest/server/helpers/router"
	"log"
)

func InitServer() {
	ctx := context.Background()
	globalDeps := dependencies.LoadGlobalDependencies(ctx)

	ginServer := gin.Default()
	internalRouter := router.NewInternalRouter(ginServer, "/api")

	handler.LoadSwaggerRoutes(ginServer)
	handler.LoadExampleRoutes(globalDeps, internalRouter)

	err := ginServer.Run(":8000")
	if err != nil {
		log.Fatalf("error to init server at PORT :8000. Error: %s", err.Error())
	}
}
