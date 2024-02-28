package service

import (
	"errors"
	"testPaymentSystem/configs"
	"testPaymentSystem/internal/domain"
	"testPaymentSystem/internal/repository"
)

type PaymentService struct {
	repository *repository.AccountRepository
	config     *configs.Config
}

func NewPaymentService(repository *repository.AccountRepository) *PaymentService {
	return &PaymentService{
		repository: repository,
	}
}

func (p *PaymentService) transactionValidation(
	accountFrom, accountTo domain.PaymentDTO,
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

func (p *PaymentService) getAccounts(
	accountNumberFrom string,
	accountNumberTo string,
) (domain.PaymentDTO, domain.PaymentDTO, bool) {
	accountFrom, ok := p.repository.GetAccount(accountNumberFrom)
	if !ok {
		return domain.PaymentDTO{}, domain.PaymentDTO{}, false
	}
	accountTo, ok := p.repository.GetAccount(accountNumberTo)
	if !ok {
		return domain.PaymentDTO{}, domain.PaymentDTO{}, false
	}
	return accountFrom, accountTo, true
}

func (p *PaymentService) Send(
	accountNumberFrom string,
	accountNumberTo string,
	sum float64,
) (bool, error) {
	// open transaction
	accountFrom, accountTo, ok := p.getAccounts(accountNumberFrom, accountNumberTo)
	if !ok {
		return false, errors.New("Account not found")
	}
	isValid, err := p.transactionValidation(accountFrom, accountTo, sum)
	if !isValid {
		return false, err
	}
	accountFrom.Balance -= sum
	accountTo.Balance += sum

	ok, err = p.repository.TransferMoney(accountFrom, accountTo)
	if err != nil {
		return false, err
	}
	// close transaction
	return ok, nil
}
