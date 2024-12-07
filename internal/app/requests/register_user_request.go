package requests

type RegisterUserRequest struct {
	Fullname string `validate:"required,max=255,min=3" json:"fullname"`
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,max=12,min=6" json:"password"`
}
