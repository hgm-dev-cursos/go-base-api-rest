package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/henriquegmendes/go-base-api-rest/dtos/request"
	"github.com/henriquegmendes/go-base-api-rest/dtos/response"
	internalErrors "github.com/henriquegmendes/go-base-api-rest/errors"
	mocks "github.com/henriquegmendes/go-base-api-rest/mocks/service"
	"github.com/henriquegmendes/go-base-api-rest/server/helpers/middlewares"
	"github.com/henriquegmendes/go-base-api-rest/server/helpers/router"
	"github.com/henriquegmendes/go-base-api-rest/service"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_exampleHandler_Create(t *testing.T) {
	type fields struct {
		exampleService func(t *testing.T) service.ExampleService
	}
	type args struct {
		body    request.ExampleRequest
		headers map[string]string
	}

	gin.SetMode(gin.ReleaseMode)

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantStatus   int
		wantResponse any
	}{
		{
			name: "Should return an error response with 400 status code due to missing Authorization header",
			fields: fields{
				exampleService: func(t *testing.T) service.ExampleService {
					m := &mocks.ExampleService{}
					return m
				},
			},
			args: args{
				body: request.ExampleRequest{
					Name: "Test Example",
				},
				headers: map[string]string{
					"Content-Type": "application/json",
				},
			},
			wantStatus: 400,
			wantResponse: internalErrors.ApplicationError{
				Message: "missing Authorization header",
				Details: []internalErrors.ErrorDetails{},
			},
		},
		{
			name: "Should return an error response with 401 status code due to invalid Authorization header",
			fields: fields{
				exampleService: func(t *testing.T) service.ExampleService {
					m := &mocks.ExampleService{}
					return m
				},
			},
			args: args{
				body: request.ExampleRequest{
					Name: "Test Example",
				},
				headers: map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "wrong-header",
				},
			},
			wantStatus: 401,
			wantResponse: internalErrors.ApplicationError{
				Message: "invalid Authorization header",
				Details: []internalErrors.ErrorDetails{},
			},
		},
		{
			name: "Should return an error response with 400 status code due to required field validation error",
			fields: fields{
				exampleService: func(t *testing.T) service.ExampleService {
					m := &mocks.ExampleService{}
					return m
				},
			},
			args: args{
				body: request.ExampleRequest{
					Name: "",
				},
				headers: map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "super-secret-auth",
				},
			},
			wantStatus: 400,
			wantResponse: internalErrors.ApplicationError{
				Message: "validation error",
				Details: []internalErrors.ErrorDetails{
					{
						Field:  "name",
						Reason: "required",
					},
				},
			},
		},
		{
			name: "Should return an error response with 400 status code due to min length field validation error",
			fields: fields{
				exampleService: func(t *testing.T) service.ExampleService {
					m := &mocks.ExampleService{}
					return m
				},
			},
			args: args{
				body: request.ExampleRequest{
					Name: "Te",
				},
				headers: map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "super-secret-auth",
				},
			},
			wantStatus: 400,
			wantResponse: internalErrors.ApplicationError{
				Message: "validation error",
				Details: []internalErrors.ErrorDetails{
					{
						Field:  "name",
						Reason: "min",
					},
				},
			},
		},
		{
			name: "Should return an internal server error response with 500 status code due to not mapped error returned from service",
			fields: fields{
				exampleService: func(t *testing.T) service.ExampleService {
					m := &mocks.ExampleService{}
					m.
						On(
							"Create",
							mock.MatchedBy(func(ctx *gin.Context) bool {
								matchAuthHeader := ctx.Request.Header.Get("Authorization") == "super-secret-auth"
								return matchAuthHeader
							}),
							request.ExampleRequest{
								Name: "Test Example",
							},
						).
						Return(nil, errors.New("not expected error")).
						Once()
					t.Cleanup(func() {
						m.AssertExpectations(t)
					})
					return m
				},
			},
			args: args{
				body: request.ExampleRequest{
					Name: "Test Example",
				},
				headers: map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "super-secret-auth",
				},
			},
			wantStatus:   500,
			wantResponse: internalErrors.DefaultApplicationError,
		},
		{
			name: "Should return an application error response with 422 status code due to error being returned from service",
			fields: fields{
				exampleService: func(t *testing.T) service.ExampleService {
					m := &mocks.ExampleService{}
					m.
						On(
							"Create",
							mock.MatchedBy(func(ctx *gin.Context) bool {
								matchAuthHeader := ctx.Request.Header.Get("Authorization") == "super-secret-auth"
								return matchAuthHeader
							}),
							request.ExampleRequest{
								Name: "Test Example",
							},
						).
						Return(nil, internalErrors.NewApplicationError("app error", http.StatusUnprocessableEntity)).
						Once()
					t.Cleanup(func() {
						m.AssertExpectations(t)
					})
					return m
				},
			},
			args: args{
				body: request.ExampleRequest{
					Name: "Test Example",
				},
				headers: map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "super-secret-auth",
				},
			},
			wantStatus: 422,
			wantResponse: internalErrors.ApplicationError{
				Message: "app error",
				Details: []internalErrors.ErrorDetails{},
			},
		},
		{
			name: "Should return an 201 response after successfully create a new example",
			fields: fields{
				exampleService: func(t *testing.T) service.ExampleService {
					m := &mocks.ExampleService{}
					m.
						On(
							"Create",
							mock.MatchedBy(func(ctx *gin.Context) bool {
								matchAuthHeader := ctx.Request.Header.Get("Authorization") == "super-secret-auth"
								return matchAuthHeader
							}),
							request.ExampleRequest{
								Name: "Test Example",
							},
						).
						Return(&response.ExampleResponse{
							ID:   "123456",
							Name: "Test Example",
						}, nil).
						Once()
					t.Cleanup(func() {
						m.AssertExpectations(t)
					})
					return m
				},
			},
			args: args{
				body: request.ExampleRequest{
					Name: "Test Example",
				},
				headers: map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "super-secret-auth",
				},
			},
			wantStatus: 201,
			wantResponse: &response.ExampleResponse{
				ID:   "123456",
				Name: "Test Example",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &exampleHandler{
				exampleService: tt.fields.exampleService(t),
			}

			ginServer := gin.New()
			internalRouter := router.NewInternalRouter(ginServer, "")
			//internalRouter.USE(middlewares.AuthMiddleware)
			internalRouter.POST("/test", middlewares.AuthRouteMiddleware(h.Create))

			bodyBytes, _ := json.Marshal(tt.args.body)
			req, _ := http.NewRequest(http.MethodPost, "/test", bytes.NewReader(bodyBytes))
			for key, value := range tt.args.headers {
				req.Header.Set(key, value)
			}

			writer := httptest.NewRecorder()

			ginServer.ServeHTTP(writer, req)

			if writer.Code != tt.wantStatus {
				t.Errorf("Create() status code = %v, wantStatus %v", writer.Code, tt.wantStatus)
			}

			wantResponseBytes, _ := json.Marshal(tt.wantResponse)
			if writer.Body.String() != string(wantResponseBytes) {
				t.Errorf("Create() response = %s, wantResponse %s", writer.Body.String(), string(wantResponseBytes))
			}
		})
	}
}
