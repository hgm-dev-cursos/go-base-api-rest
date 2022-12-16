package request

type ExampleRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}
