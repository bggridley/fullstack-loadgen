package request

type CreateTestRequest struct {
	Name string `validate:"required,min=1,max=16" json:"name"`
}