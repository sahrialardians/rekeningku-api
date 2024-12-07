package services

import (
	"errors"

	"github.com/sahrialardians/rekeningku/internal/app/requests"
	"github.com/sahrialardians/rekeningku/internal/models"
	"github.com/sahrialardians/rekeningku/internal/repositories"
	"github.com/sahrialardians/rekeningku/internal/utils"
)

type AuthServiceImpl struct {
	userRepository repositories.UserRepository
}

func NewAuthServiceImpl(userRepo repositories.UserRepository) AuthService {
	return &AuthServiceImpl{
		userRepository: userRepo,
	}
}

func (authServiceImpl *AuthServiceImpl) Register(req requests.RegisterUserRequest) (string, error) {
	// Hash password
	hashedPassword, err := utils.HashedPassword(req.Password)
	if err != nil {
		return "", err
	}

	// Create user
	user := models.User{
		Fullname: req.Fullname,
		Email:    req.Email,
		Password: hashedPassword,
	}

	// Insert to user table
	savedUser, err := authServiceImpl.userRepository.Save(&user)
	if err != nil {
		return "", err
	}

	// Generate JWT
	token, err := utils.GenerateJWT(savedUser.ID, savedUser.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (authServiceImpl *AuthServiceImpl) Login(req requests.LoginUserRequest) (string, error) {
	// Find user by email
	user, err := authServiceImpl.userRepository.FindByEmail(req.Email)
	if err != nil {
		return "", errors.New("email or password is wrong")
	}

	// Validate password
	if !utils.VerifyPassword(user.Password, req.Password) {
		return "", errors.New("email or password is wrong")
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
