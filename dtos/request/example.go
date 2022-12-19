package request

// CreateExampleRequest
// swagger:parameters CreateExampleRequest
type CreateExampleRequest struct {
	AuthorizationHeader

	// Request to create a new example
	//
	// required: true
	// in: body
	CreateExampleRequest ExampleRequest
}

type AuthorizationHeader struct {
	// Authorization header
	//
	// required: true
	// in: header
	Authorization string `json:"Authorization"`
}

type ExampleRequest struct {
	// the name for this user
	Name string `json:"name" validate:"required,min=3"`
}
