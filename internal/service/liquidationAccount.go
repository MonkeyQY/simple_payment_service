package service

import "testPaymentSystem/internal/repository"

type LiquidationSpecialAccountService struct {
	repository *repository.AccountRepository
}

func (s *LiquidationSpecialAccountService) GetAccountNumber() string {
	return "BY04CBDC36029110100040000001"
}

func (s *LiquidationSpecialAccountService) Add(sum float64) (float64, error) {
	liquidationAccount, ok := s.repository.GetAccount("BY04CBDC36029110100040000001")
	if !ok {
		panic("Liquidation account not found")
	}
	liquidationAccount.Balance += sum
	err := s.repository.AddAccount(liquidationAccount)
	if err != nil {
		return 0, err
	}
	return sum, err
}

func NewLiquidationSpecialAccountService(
	repository *repository.AccountRepository,
) *LiquidationSpecialAccountService {
	return &LiquidationSpecialAccountService{
		repository: repository,
	}
}
