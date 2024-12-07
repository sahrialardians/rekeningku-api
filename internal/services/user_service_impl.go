package services

import (
	"github.com/sahrialardians/rekeningku/internal/app/requests"
	"github.com/sahrialardians/rekeningku/internal/app/responses"
	"github.com/sahrialardians/rekeningku/internal/repositories"
)

type UserServiceImpl struct {
	UserRepository repositories.UserRepository
}

// NewUserServiceImpl creates a new instance of UserServiceImpl
func NewUserServiceImpl(userRepo repositories.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: userRepo,
	}
}

// Update updates an existing user's profile
func (userServiceImpl UserServiceImpl) Update(user requests.UpdateUserRequest) error {
	// Find the user by ID
	userData, err := userServiceImpl.UserRepository.FindById(user.ID)
	if err != nil {
		return err
	}

	// Update user data
	userData.Fullname = user.Fullname
	return userServiceImpl.UserRepository.Update(&userData)
}

// Delete removes a user by ID
func (userServiceImpl UserServiceImpl) Delete(usersId int) error {

	err := userServiceImpl.UserRepository.Delete(usersId)
	if err != nil {
		return err
	}

	return nil
}

// FindById retrieves a user by ID
func (userServiceImpl UserServiceImpl) FindById(usersId int) (responses.UserResponse, error) {
	// Find the user
	userData, err := userServiceImpl.UserRepository.FindById(usersId)
	if err != nil {
		return responses.UserResponse{}, err
	}

	// Map the model to the response
	userResponse := responses.UserResponse{
		ID:       userData.ID,
		Fullname: userData.Fullname,
		Email:    userData.Email,
	}

	return userResponse, nil
}
