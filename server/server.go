package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/henriquegmendes/go-base-api-rest/client"
	"github.com/henriquegmendes/go-base-api-rest/dtos/request"
	"github.com/henriquegmendes/go-base-api-rest/handler"
	"github.com/henriquegmendes/go-base-api-rest/repository"
	"github.com/henriquegmendes/go-base-api-rest/server/helpers/router"
	"github.com/henriquegmendes/go-base-api-rest/service"
	"log"
)

func InitServer() {
	ginServer := gin.Default()
	internalRouter := router.NewInternalRouter(ginServer, "/api")

	ctx := context.Background()
	mongoDep, _ := client.NewMongoClient(ctx)
	exampleRepository := repository.NewExampleRepository(mongoDep.ExampleDatabase)
	exampleService := service.NewExampleService(exampleRepository)
	response, err := exampleService.Create(ctx, request.ExampleRequest{
		Name: "New test example from service",
	})

	log.Printf("Response from service ---> %v / Error: %v", response, err)

	handler.LoadExampleRoutes(internalRouter)

	err = ginServer.Run(":8000")
	if err != nil {
		log.Fatalf("error to init server at PORT :8000. Error: %s", err.Error())
	}
}
