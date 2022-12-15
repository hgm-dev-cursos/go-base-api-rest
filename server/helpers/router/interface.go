package router

import "github.com/gin-gonic/gin"

type InternalRouter interface {
	POST(path string, handler InternalHandler)
}

type InternalHandler func(ctx *gin.Context) (*InternalResponse, error)

type InternalResponse struct {
	Body       any
	StatusCode int
}

func NewInternalResponse(body any, statusCode int) *InternalResponse {
	return &InternalResponse{
		Body:       body,
		StatusCode: statusCode,
	}
}
