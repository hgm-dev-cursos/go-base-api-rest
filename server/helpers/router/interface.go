package router

import "github.com/gin-gonic/gin"

type InternalRouter interface {
	POST(path string, handler InternalHandler)
	USE(handler InternalMiddlewareHandler)
}

type InternalHandler func(ctx *gin.Context) (*InternalResponse, error)
type InternalMiddlewareHandler func(ctx *gin.Context) error

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
