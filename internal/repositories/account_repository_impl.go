package repositories

import (
	"errors"

	"github.com/sahrialardians/rekeningku/internal/app/requests"
	"github.com/sahrialardians/rekeningku/internal/models"
	"gorm.io/gorm"
)

type AccountRepositoryImpl struct {
	Db *gorm.DB
}

// Contract AccountRepository
func NewAccountRepositoryImpl(Db *gorm.DB) AccountRepository {
	return &AccountRepositoryImpl{Db: Db}
}

// CountAll implements AccountRepository
func (repository *AccountRepositoryImpl) CountAll(userId int) (int64, error) {
	var totalRecords int64

	err := repository.Db.Model(&models.Account{}).Where("user_id = ?", userId).Count(&totalRecords).Error
	if err != nil {
		return 0, err
	}

	return totalRecords, nil
}

// FindAll implement AccountRepository
func (repository *AccountRepositoryImpl) FindAll(userId, page, pageSize int) ([]models.Account, int, error) {
	var accounts []models.Account

	// Calculate offset
	offset := (page - 1) * pageSize

	// Fetch paginated records
	err := repository.Db.Preload("User").Where("user_id = ?", userId).Limit(pageSize).Offset(offset).Find(&accounts).Error
	if err != nil {
		return nil, 0, err
	}

	return accounts, len(accounts), nil
}

// Save implements AccountRepository.
func (repository AccountRepositoryImpl) Save(account *models.Account) (*models.Account, error) {
	result := repository.Db.Create(account)
	if result.Error != nil {
		return nil, result.Error
	}

	return account, nil
}

// Update implements AccountRepository.
func (repository AccountRepositoryImpl) Update(accounts *models.Account) error {
	var updateaccount = requests.UpdateAccountRequest{
		ID:            accounts.ID,
		AccountName:   accounts.AccountName,
		AccountCode:   accounts.AccountCode,
		AccountNumber: accounts.AccountNumber,
	}

	result := repository.Db.Model(&accounts).Where("id = ? AND user_id = ?", accounts.ID, accounts.UserID).Updates(updateaccount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindById implements AccountRepository.
func (repository AccountRepositoryImpl) FindById(userId, accountsId int) (models.Account, error) {
	var account models.Account
	result := repository.Db.Where("user_id = ?", userId).First(&account, accountsId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return account, errors.New("account is not found")
		}
		return account, result.Error
	}
	return account, nil
}

// Delete implements AccountRepository.
func (repository AccountRepositoryImpl) Delete(userId, accountsId int) error {
	var accounts models.Account
	result := repository.Db.Where("id = ? AND user_id = ?", accountsId, userId).Delete(&accounts)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
