package domain

type Account struct {
	AccountNumber string
	Balance       float64
	Active        bool
	CreatedAt     string
	UpdatedAt     string
	Currency      string
	Limits        bool
	Special       bool
}
