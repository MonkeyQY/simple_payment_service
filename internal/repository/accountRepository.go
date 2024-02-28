package repository

import (
	"testPaymentSystem/internal/db"
	"testPaymentSystem/internal/domain"
)

// Error может быть вызвана в реальном приложении при работе с реальной базой данных
// поэтому важно возвращать ошибку, чтобы обработать ее в реальном приложении
type AccountRepository struct {
	Db *db.DB
}

func NewAccountRepository(db *db.DB) *AccountRepository {
	return &AccountRepository{
		Db: db,
	}
}

func (d *AccountRepository) AddAccount(account domain.PaymentDTO) error {
	d.Db.Accounts[account.AccountNumber] = account
	return nil
}

func (d *AccountRepository) GetAccount(accountNumber string) (domain.PaymentDTO, bool) {
	account, ok := d.Db.Accounts[accountNumber]
	if !ok {
		return domain.PaymentDTO{}, false
	}
	return account, true
}

func (d *AccountRepository) GetAccounts() []domain.PaymentDTO {
	var accounts []domain.PaymentDTO
	for _, account := range d.Db.Accounts {
		accounts = append(accounts, account)
	}
	return accounts
}

func (d *AccountRepository) TransferMoney(
	accountFrom, accountTo domain.PaymentDTO,
) (bool, error) {
	// В Реальной базе, мы откроем транзакцию и если одна из операций не прошла, то откатим транзакцию
	// Работает это по-другому, но нужно понимать, что важно обработать ошибку
	oldValueFrom := d.Db.Accounts[accountFrom.AccountNumber]
	oldValueTo := d.Db.Accounts[accountTo.AccountNumber]

	d.Db.Accounts[accountFrom.AccountNumber] = accountFrom
	k := d.Db.Accounts[accountTo.AccountNumber]
	if k == (oldValueFrom) {
		return false, nil
	}

	d.Db.Accounts[accountTo.AccountNumber] = accountTo
	v := d.Db.Accounts[accountTo.AccountNumber]
	if v == (oldValueTo) {
		d.Db.Accounts[accountFrom.AccountNumber] = oldValueFrom
		return false, nil
	}

	return true, nil
}