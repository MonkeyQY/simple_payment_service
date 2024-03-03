package repository

import (
	"testPaymentSystem/internal/db"
	"testPaymentSystem/internal/domain"
	"time"
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

// AddAccount - добавление счета в базу данных. в релаьной базе updated_at будут занесены автоматически
func (d *AccountRepository) AddAccount(account domain.Account) error {
	account.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	d.Db.Accounts[account.AccountNumber] = account
	return nil
}

func (d *AccountRepository) GetAccount(accountNumber string) (domain.Account, bool) {
	account, ok := d.Db.Accounts[accountNumber]
	if !ok {
		return domain.Account{}, false
	}
	return account, true
}

func (d *AccountRepository) GetAccounts() []domain.Account {
	var accounts []domain.Account
	for _, account := range d.Db.Accounts {
		accounts = append(accounts, account)
	}
	return accounts
}

func (d *AccountRepository) TransferMoney(
	accountFrom, accountTo domain.Account,
) (bool, error) {
	// В Реальной базе, мы откроем транзакцию и если одна из операций не прошла, то откатим транзакцию

	accountFrom.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	accountTo.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	d.Db.Accounts[accountFrom.AccountNumber] = accountFrom

	d.Db.Accounts[accountTo.AccountNumber] = accountTo

	return true, nil
}
