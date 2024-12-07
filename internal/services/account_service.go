package services

import (
	"github.com/sahrialardians/rekeningku/internal/app/requests"
	"github.com/sahrialardians/rekeningku/internal/app/responses"
)

type AccountService interface {
	FindAll(userId, page, pageSize int) ([]responses.AccountResponse, int, error)
	FindById(userId, accountsId int) (responses.AccountResponse, error)
	Save(account requests.CreateAccountRequest) (responses.AccountResponse, error)
	Update(account requests.UpdateAccountRequest) error
	Delete(userId, accountsId int) error
}
