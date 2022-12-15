package server

import (
	"github.com/gin-gonic/gin"
	"github.com/henriquegmendes/go-base-api-rest/handler"
	"log"
)

func InitServer() {
	ginServer := gin.Default()

	handler.LoadExampleRoutes(ginServer)

	err := ginServer.Run(":8000")
	if err != nil {
		log.Fatalf("error to init server at PORT :8000. Error: %s", err.Error())
	}
}
