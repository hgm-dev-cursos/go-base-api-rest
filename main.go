package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	ginServer := gin.Default()

	ginServer.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Hello World",
		})
	})

	err := ginServer.Run(":8000")
	if err != nil {
		log.Fatalf("error to init server at PORT :8000. Error: %s", err.Error())
	}
}
