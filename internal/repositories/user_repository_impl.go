package repositories

import (
	"errors"

	"github.com/sahrialardians/rekeningku/internal/app/requests"
	"github.com/sahrialardians/rekeningku/internal/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

// Contract UserRepository
func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Save implements UserRepository.
func (userRepoImpl UserRepositoryImpl) Save(user *models.User) (*models.User, error) {
	result := userRepoImpl.Db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// Update implements UserRepository.
func (userRepoImpl UserRepositoryImpl) Update(users *models.User) error {
	var updateUser = requests.UpdateUserRequest{
		ID:       users.ID,
		Fullname: users.Fullname,
	}

	result := userRepoImpl.Db.Model(&users).Updates(updateUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindById implements UserRepository.
func (userRepoImpl UserRepositoryImpl) FindById(usersId int) (models.User, error) {
	var user models.User
	result := userRepoImpl.Db.Where("id= ?", usersId).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("user is not found")
		}
		return user, result.Error
	}
	return user, nil
}

// FindByEmail implements UserRepository.
func (userRepoImpl *UserRepositoryImpl) FindByEmail(email string) (models.User, error) {
	var user models.User
	result := userRepoImpl.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, result.Error
	}
	return user, nil
}

// Delete implements UserRepository.
func (userRepoImpl UserRepositoryImpl) Delete(usersId int) error {
	var users models.User
	result := userRepoImpl.Db.Where("id = ?", usersId).Delete(&users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
