package errors

type ApplicationError struct {
	StatusCode int            `json:"-"`
	Message    string         `json:"message"`
	Details    []ErrorDetails `json:"details"`
}

type ErrorDetails struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
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
