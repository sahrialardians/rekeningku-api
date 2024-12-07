package repositories

import "github.com/sahrialardians/rekeningku/internal/models"

type AccountRepository interface {
	CountAll(userId int) (int64, error)
	FindAll(userId, page, pageSize int) ([]models.Account, int, error)
	FindById(userId, accountsId int) (models.Account, error)
	Save(account *models.Account) (*models.Account, error)
	Update(account *models.Account) error
	Delete(userId, accountsId int) error
}
