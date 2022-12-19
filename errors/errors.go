package errors

import "github.com/henriquegmendes/go-base-api-rest/dtos/response"

type ApplicationError struct {
	StatusCode int
	Message    string
	Details    []ErrorDetails
}

type ErrorDetails struct {
	Field  string
	Reason string
}

func (e *ApplicationError) ToErrorResponse() response.ErrorResponse {
	errorDetails := make([]response.ErrorDetailsResponse, 0)
	if e.Details != nil && len(e.Details) > 0 {
		for _, detail := range e.Details {
			errorDetail := response.ErrorDetailsResponse{
				Field:  detail.Field,
				Reason: detail.Reason,
			}
			errorDetails = append(errorDetails, errorDetail)
		}
	}

	return response.ErrorResponse{
		Message: e.Message,
		Details: errorDetails,
	}
}

func (e *ApplicationError) Error() string {
	return e.Message
}

func NewApplicationError(message string, statusCode int) *ApplicationError {
	return &ApplicationError{
		StatusCode: statusCode,
		Message:    message,
		Details:    []ErrorDetails{},
	}
}

var (
	DefaultApplicationError = NewApplicationError("internal server error", 500)
)
