package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/henriquegmendes/go-base-api-rest/cfg"
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

	port := fmt.Sprintf(":%v", cfg.Env().Port)
	err := ginServer.Run(port)
	if err != nil {
		log.Fatalf("error to init server at PORT %s. Error: %s", port, err.Error())
	}
}
