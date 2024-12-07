package services

import (
	"github.com/sahrialardians/rekeningku/internal/app/requests"
	"github.com/sahrialardians/rekeningku/internal/app/responses"
	"github.com/sahrialardians/rekeningku/internal/models"
	"github.com/sahrialardians/rekeningku/internal/repositories"
)

type AccountServiceImpl struct {
	accountRepository repositories.AccountRepository
}

// NewAccountServiceImpl creates a new instance of AccountServiceImpl
func NewAccountServiceImpl(accountRepository repositories.AccountRepository) AccountService {
	return &AccountServiceImpl{
		accountRepository: accountRepository,
	}
}

// FindAll implement AccountService
func (service *AccountServiceImpl) FindAll(userId, page, pageSize int) ([]responses.AccountResponse, int, error) {
	// Count total records
	totalRecords, err := service.accountRepository.CountAll(userId)
	if err != nil {
		return nil, 0, err
	}

	// Fetch paginated records
	result, _, err := service.accountRepository.FindAll(userId, page, pageSize)

	var accounts []responses.AccountResponse
	if err != nil {
		return nil, 0, err
	}

	for _, value := range result {
		account := responses.AccountResponse{
			ID:                value.ID,
			AccountName:       value.AccountName,
			AccountCode:       value.AccountCode,
			AccountNumber:     value.AccountNumber,
			AccountHolderName: value.AccountHolderName,
			CreatedAt:         value.CreatedAt,
			UpdatedAt:         value.UpdatedAt,
		}
		accounts = append(accounts, account)
	}

	return accounts, int(totalRecords), nil
}

// FindById implements AccountService.
func (service *AccountServiceImpl) FindById(userId, accountsId int) (responses.AccountResponse, error) {
	// Find the account
	account, err := service.accountRepository.FindById(userId, accountsId)
	if err != nil {
		return responses.AccountResponse{}, err
	}

	// Map the model to the response
	result := responses.AccountResponse{
		ID:                account.ID,
		AccountName:       account.AccountName,
		AccountNumber:     account.AccountNumber,
		AccountCode:       account.AccountCode,
		AccountHolderName: account.AccountHolderName,
	}

	return result, nil
}

// Save implements AccountService.
func (service *AccountServiceImpl) Save(req requests.CreateAccountRequest) (responses.AccountResponse, error) {

	// Map request to account model
	account := models.Account{
		UserID:            int(req.UserID),
		AccountName:       req.AccountName,
		AccountCode:       req.AccountCode,
		AccountNumber:     req.AccountNumber,
		AccountHolderName: req.AccountHolderName,
	}

	// Save account to repository
	savedAccount, err := service.accountRepository.Save(&account)
	if err != nil {
		return responses.AccountResponse{}, err
	}

	// Map saved account to response
	accountResponse := responses.AccountResponse{
		ID:                savedAccount.ID,
		AccountName:       savedAccount.AccountName,
		AccountCode:       savedAccount.AccountCode,
		AccountNumber:     savedAccount.AccountNumber,
		AccountHolderName: savedAccount.AccountHolderName,
		CreatedAt:         savedAccount.CreatedAt,
		UpdatedAt:         savedAccount.UpdatedAt,
	}

	return accountResponse, nil
}

// Update implements AccountService.
func (service *AccountServiceImpl) Update(account requests.UpdateAccountRequest) error {
	// Find the account by ID
	accountData, err := service.accountRepository.FindById(account.UserID, account.ID)
	if err != nil {
		return err
	}

	// Update account data
	accountData.AccountName = account.AccountName
	accountData.AccountCode = account.AccountCode
	accountData.AccountNumber = account.AccountNumber

	err = service.accountRepository.Update(&accountData)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements AccountService.
func (service *AccountServiceImpl) Delete(userId, accountsId int) error {
	err := service.accountRepository.Delete(userId, accountsId)
	if err != nil {
		return err
	}
	return nil
}
