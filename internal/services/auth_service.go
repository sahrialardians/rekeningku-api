package services

import (
	"github.com/sahrialardians/rekeningku/internal/app/requests"
)

type AuthService interface {
	Register(requests.RegisterUserRequest) (string, error)
	Login(requests.LoginUserRequest) (string, error)
}
