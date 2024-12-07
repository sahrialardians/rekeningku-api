package repositories

import "github.com/sahrialardians/rekeningku/internal/models"

type UserRepository interface {
	Save(user *models.User) (*models.User, error)
	Update(user *models.User) error
	Delete(usersId int) error
	FindById(usersId int) (models.User, error)
	FindByEmail(email string) (models.User, error)
}
