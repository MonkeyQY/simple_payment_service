package service

import "testPaymentSystem/internal/repository"

type EmissionSpecialAccountService struct {
	repository *repository.AccountRepository
}

func NewEmissionSpecialAccountService(
	repository *repository.AccountRepository,
) *EmissionSpecialAccountService {
	return &EmissionSpecialAccountService{
		repository: repository,
	}
}

func (s *EmissionSpecialAccountService) GetAccountNumber() string {
	return "BY04CBDC36029110100040000000"
}

func (s *EmissionSpecialAccountService) Add(sum float64) (float64, error) {
	nationAccount, ok := s.repository.GetAccount("BY04CBDC36029110100040000000")
	if !ok {
		panic("Nation account not found")
	}
	nationAccount.Balance += sum
	err := s.repository.AddAccount(nationAccount)
	if err != nil {
		return 0, err
	}
	return sum, err
}
