package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadExampleRoutes(ginServer *gin.Engine) {
	ginServer.POST("/api/example", Create)
}

func Create(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Hello World",
	})
}
