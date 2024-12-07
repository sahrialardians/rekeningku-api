package requests

type UpdateUserRequest struct {
	ID       int    `validate:"required"`
	Fullname string `validate:"required,max=255,min=3" json:"fullname"`
}
