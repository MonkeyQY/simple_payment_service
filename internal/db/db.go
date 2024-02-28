package db

import (
	"sync"
	"testPaymentSystem/configs"
	"testPaymentSystem/internal/domain"
)

var (
	Instance *DB
	once     sync.Once
)

// Imitation of the DB
type DB struct {
	Accounts map[string]domain.PaymentDTO
}

func NewDB(config *configs.Config) *DB {
	// In real application, we will connect to the database
	// Once , чтобы не создавать несколько экземпляров и всегда обращаться к одной мок базе
	once.Do(func() {
		Instance = &DB{
			Accounts: make(map[string]domain.PaymentDTO),
		}
	})
	return Instance
}
