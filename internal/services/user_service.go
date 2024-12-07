package services

import (
	"github.com/sahrialardians/rekeningku/internal/app/requests"
	"github.com/sahrialardians/rekeningku/internal/app/responses"
)

type UserService interface {
	Update(user requests.UpdateUserRequest) error
	Delete(usersId int) error
	FindById(usersId int) (responses.UserResponse, error)
}
