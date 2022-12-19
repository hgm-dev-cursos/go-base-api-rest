package response

// CreateExampleResponse
// swagger:response CreateExampleResponse
type CreateExampleResponse struct {
	// in: body
	Body ExampleResponse
}

type ExampleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ApplicationErrorResponse
// swagger:response ApplicationErrorResponse
type ApplicationErrorResponse struct {
	// in: body
	Body ErrorResponse
}

type ErrorResponse struct {
	Message string                 `json:"message"`
	Details []ErrorDetailsResponse `json:"details"`
}

type ErrorDetailsResponse struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}
