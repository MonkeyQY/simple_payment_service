package service

import "testPaymentSystem/internal/domain"

// Account обьявляет методы, который должны быть для того, чтобы работать с Счетами пользователей
type Account interface {
	NewAccount() (domain.PaymentDTO, error)
	GetAccounts() []domain.PaymentDTO
	Replenishment(accountNumber string, sum float64) (float64, error)
}

type SpecialAccount interface {
	GetAccountNumber() string
	Add(sum float64) (float64, error)
}

type Payment interface {
	Send(
		accountNumberFrom string,
		accountNumberTo string,
		sum float64,
	) (bool, error)
}

type PaymentSystem struct {
	emissionSpecialAccountService    SpecialAccount
	liquidationSpecialAccountService SpecialAccount
	accountService                   Account
	paymentService                   Payment
}

func NewPaymentSystem(
	emissionSpecialAccountService SpecialAccount,
	liquidationSpecialAccountService SpecialAccount,
	accountService Account,
	paymentService Payment,
) *PaymentSystem {
	return &PaymentSystem{
		emissionSpecialAccountService:    emissionSpecialAccountService,
		liquidationSpecialAccountService: liquidationSpecialAccountService,
		accountService:                   accountService,
		paymentService:                   paymentService,
	}
}

func (p *PaymentSystem) GetNationAccountNumber() string {
	return p.emissionSpecialAccountService.GetAccountNumber()
}

func (p *PaymentSystem) GetLiquidationAccountNumber() string {
	return p.liquidationSpecialAccountService.GetAccountNumber()
}

func (p *PaymentSystem) AddToNationAccount(sum float64) (float64, error) {
	return p.emissionSpecialAccountService.Add(sum)
}

func (p *PaymentSystem) AddToLiquidationAccount(sum float64) (float64, error) {
	return p.liquidationSpecialAccountService.Add(sum)
}

func (p *PaymentSystem) Transfer(
	accountNumberFrom string,
	accountNumberTo string,
	sum float64,
) (bool, error) {
	ok, err := p.paymentService.Send(accountNumberFrom, accountNumberTo, sum)
	if err != nil {
		return false, err
	}
	return ok, nil
}

// TransferWithMap - метод для трансфера с использованием json
func (p *PaymentSystem) TransferWithMap(data map[string]interface{}) (bool, error) {
	accountNumberFrom, ok := data["accountNumberFrom"].(string)
	if !ok {
		return false, nil
	}
	accountNumberTo, ok := data["accountNumberTo"].(string)
	if !ok {
		return false, nil
	}
	sum, ok := data["sum"].(float64)
	if !ok {
		return false, nil
	}
	return p.Transfer(accountNumberFrom, accountNumberTo, sum)
}

func (p *PaymentSystem) GetAllAccounts() []domain.PaymentDTO {
	return p.accountService.GetAccounts()
}

func (p *PaymentSystem) CreateNewAccount() (string, error) {
	account, err := p.accountService.NewAccount()
	if err != nil {
		return "", err
	}
	return account.AccountNumber, nil
}

func (p *PaymentSystem) Replenishment(accountNumber string, sum float64) (float64, error) {
	return p.accountService.Replenishment(accountNumber, sum)
}
