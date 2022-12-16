package helpers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	internalErrors "github.com/henriquegmendes/go-base-api-rest/errors"
	"net/http"
)

var validate = validator.New()

func UnmarshalAndValidate(dataBytes []byte, request any) error {
	err := json.Unmarshal(dataBytes, request)
	if err != nil {
		return internalErrors.NewApplicationError("invalid json body request", http.StatusBadRequest)
	}

	err = validate.Struct(request)
	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)
	applicationError := internalErrors.NewApplicationError("validation error", http.StatusBadRequest)

	for _, validationErr := range validationErrors {
		applicationError.Details = append(applicationError.Details, internalErrors.ErrorDetails{
			Field:  NormalizeFieldNameByTag(request, validationErr.StructField(), "json"),
			Reason: validationErr.Tag(),
		})
	}

	return applicationError
}
