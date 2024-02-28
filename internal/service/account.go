package service

import (
	"math/rand"
	"strconv"
	"testPaymentSystem/internal/domain"
	"testPaymentSystem/internal/repository"
)

type AccountService struct {
	repository *repository.AccountRepository
}

func NewAccountService(repository *repository.AccountRepository) *AccountService {
	return &AccountService{
		repository: repository,
	}
}

func (a *AccountService) generateAccountNumber() string {
	nationPrefix := "BY09"
	codeBank := "CBDC"
	sortCode := 709080

	minR := 10000000
	maxR := 99999999
	accountNumber := rand.Intn(maxR-minR) + minR
	return nationPrefix + codeBank + strconv.Itoa(sortCode) + strconv.Itoa(accountNumber)
}

func (a *AccountService) checkAccountExist(accountNumber string) bool {
	_, ok := a.repository.GetAccount(accountNumber)
	return ok
}

func (a *AccountService) NewAccount() (domain.PaymentDTO, error) {
	accountNumber := a.generateAccountNumber()
	isExist := a.checkAccountExist(accountNumber)
	if isExist {
		return a.NewAccount()
	}
	account := domain.PaymentDTO{
		AccountNumber: accountNumber,
	}
	err := a.repository.AddAccount(account)
	if err != nil {
		return domain.PaymentDTO{}, err
	}
	return account, nil
}

func (a *AccountService) GetAccounts() []domain.PaymentDTO {
	return a.repository.GetAccounts()
}
