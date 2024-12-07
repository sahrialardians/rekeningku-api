package requests

type LoginUserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,max=12,min=6" json:"password"`
}
