package service

import (
	"math/rand"
	"strconv"
	"testPaymentSystem/internal/domain"
	"time"
)

type AccountRepository interface {
	AddAccount(account domain.Account) error
	GetAccount(accountNumber string) (domain.Account, bool)
	GetAccounts() []domain.Account
}

type AccountService struct {
	repository AccountRepository
}

func NewAccountService(repository AccountRepository) *AccountService {
	return &AccountService{
		repository: repository,
	}
}

func (a *AccountService) generateAccountNumber() string {
	nationPrefix := "BY09"
	codeBank := "CBDC"
	sortCode := 709080
	solt := "100500"

	minR := 10000000
	maxR := 99999999
	accountNumber := rand.Intn(maxR-minR) + minR
	return nationPrefix + codeBank + solt + strconv.Itoa(sortCode) + strconv.Itoa(accountNumber)
}

func (a *AccountService) checkAccountExist(accountNumber string) bool {
	_, ok := a.repository.GetAccount(accountNumber)
	return ok
}

func (a *AccountService) NewAccount() (domain.Account, error) {
	accountNumber := a.generateAccountNumber()
	isExist := a.checkAccountExist(accountNumber)
	if isExist {
		return a.NewAccount()
	}
	account := domain.Account{
		AccountNumber: accountNumber,
		Balance:       0,
		Active:        true,
		Currency:      "BYN",
		Limits:        false,
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
	err := a.repository.AddAccount(account)
	if err != nil {
		return domain.Account{}, err
	}
	return account, nil
}

func (a *AccountService) GetAccounts() []domain.Account {
	return a.repository.GetAccounts()
}

func (a *AccountService) Replenishment(accountNumber string, sum float64) (float64, error) {
	account, ok := a.repository.GetAccount(accountNumber)
	if !ok {
		return 0, nil
	}
	account.Balance += sum
	err := a.repository.AddAccount(account)
	if err != nil {
		return 0, err
	}
	return account.Balance, nil
}
