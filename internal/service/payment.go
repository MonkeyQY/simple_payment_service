package service

import (
	"errors"
	"testPaymentSystem/internal/domain"
)

type PaymentRepository interface {
	TransferMoney(accountFrom domain.Account, accountTo domain.Account) (bool, error)
	GetAccount(accountNumber string) (domain.Account, bool)
}

type PaymentService struct {
	repository PaymentRepository
}

func NewPaymentService(repository PaymentRepository) *PaymentService {
	return &PaymentService{
		repository: repository,
	}
}

// transactionValidation - валидация данных перед отправкой денег
func (p *PaymentService) transactionValidation(
	accountFrom, accountTo domain.Account,
	sum float64,
) (bool, error) {

	if accountFrom.Balance < sum {
		return false, errors.New("Not enough money")
	}
	if accountFrom.AccountNumber == accountTo.AccountNumber {
		return false, errors.New("The same account")
	}

	if !accountFrom.Active {
		return false, errors.New("The account is not active")
	}

	if accountFrom.Special {
		return false, errors.New("The account is Special")
	}
	return true, nil
}

// getAccounts - получение счетов для отправки денег по их номерам
func (p *PaymentService) getAccounts(
	accountNumberFrom string,
	accountNumberTo string,
) (domain.Account, domain.Account, bool) {
	accountFrom, ok := p.repository.GetAccount(accountNumberFrom)
	if !ok {
		return domain.Account{}, domain.Account{}, false
	}
	accountTo, ok := p.repository.GetAccount(accountNumberTo)
	if !ok {
		return domain.Account{}, domain.Account{}, false
	}
	return accountFrom, accountTo, true
}

// Send - отправка денег. Валидация, открытие транзакции, иммитация запроса в бд и  взакрытие транзакции
func (p *PaymentService) Send(
	accountNumberFrom string,
	accountNumberTo string,
	sum float64,
) (bool, error) {
	accountFrom, accountTo, ok := p.getAccounts(accountNumberFrom, accountNumberTo)
	if !ok {
		return false, errors.New("Account not found")
	}
	isValid, err := p.transactionValidation(accountFrom, accountTo, sum)
	if !isValid {
		return false, err
	}
	// open transaction
	accountFrom.Balance -= sum
	accountTo.Balance += sum

	ok, err = p.repository.TransferMoney(accountFrom, accountTo)
	if err != nil {
		return false, err
	}
	// commit or rollback
	// close transaction
	return ok, nil
}
