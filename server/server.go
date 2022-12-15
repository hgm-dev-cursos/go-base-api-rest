package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/henriquegmendes/go-base-api-rest/client"
	"github.com/henriquegmendes/go-base-api-rest/handler"
	"github.com/henriquegmendes/go-base-api-rest/server/helpers/router"
	"log"
)

func InitServer() {
	ginServer := gin.Default()
	internalRouter := router.NewInternalRouter(ginServer, "/api")

	ctx := context.Background()
	_, _ = client.NewMongoClient(ctx)

	handler.LoadExampleRoutes(internalRouter)

	err := ginServer.Run(":8000")
	if err != nil {
		log.Fatalf("error to init server at PORT :8000. Error: %s", err.Error())
	}
}
