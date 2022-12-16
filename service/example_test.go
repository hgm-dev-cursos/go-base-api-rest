package service

import (
	"context"
	"errors"
	"github.com/henriquegmendes/go-base-api-rest/dtos/request"
	mocks "github.com/henriquegmendes/go-base-api-rest/mocks/repository"
	"github.com/henriquegmendes/go-base-api-rest/models"
	"github.com/henriquegmendes/go-base-api-rest/repository"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

func Test_exampleService_Create(t *testing.T) {
	type fields struct {
		exampleRepository func(t *testing.T) repository.ExampleRepository
	}
	type args struct {
		ctx            context.Context
		exampleRequest request.ExampleRequest
	}
	type exampleResponseToValidate struct {
		Name string
	}

	mockContext := context.Background()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *exampleResponseToValidate
		wantErr bool
	}{
		{
			name: "Should return error due to error in saving new example in database",
			fields: fields{
				exampleRepository: func(t *testing.T) repository.ExampleRepository {
					m := &mocks.ExampleRepository{}
					m.
						On("Create", mockContext, mock.MatchedBy(func(arg *models.Example) bool {
							return arg.Name == "Test Example"
						})).
						Return(errors.New("error")).
						Once()
					t.Cleanup(func() {
						m.AssertExpectations(t)
					})
					return m
				},
			},
			args: args{
				ctx: mockContext,
				exampleRequest: request.ExampleRequest{
					Name: "Test Example",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Should create a new example with no errors and return a response",
			fields: fields{
				exampleRepository: func(t *testing.T) repository.ExampleRepository {
					m := &mocks.ExampleRepository{}
					m.
						On("Create", mockContext, mock.MatchedBy(func(arg *models.Example) bool {
							return arg.Name == "Test Example"
						})).
						Return(nil).
						Once()
					t.Cleanup(func() {
						m.AssertExpectations(t)
					})
					return m
				},
			},
			args: args{
				ctx: mockContext,
				exampleRequest: request.ExampleRequest{
					Name: "Test Example",
				},
			},
			want: &exampleResponseToValidate{
				Name: "Test Example",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &exampleService{
				exampleRepository: tt.fields.exampleRepository(t),
			}
			got, err := s.Create(tt.args.ctx, tt.args.exampleRequest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var gotToValidate *exampleResponseToValidate
			if got != nil {
				gotToValidate = &exampleResponseToValidate{
					Name: got.Name,
				}
			}
			if !reflect.DeepEqual(gotToValidate, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
