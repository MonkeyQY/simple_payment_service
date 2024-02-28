package app

import (
	"fmt"
	"testPaymentSystem/configs"
	"testPaymentSystem/internal/db"
	"testPaymentSystem/internal/domain"
	"testPaymentSystem/internal/repository"
	"testPaymentSystem/internal/service"
)

func CreateSpecialAccounts(accountRepository *repository.AccountRepository) error {
	nationAccount := domain.PaymentDTO{Special: true}
	_, ok := accountRepository.GetAccount(nationAccount.AccountNumber)
	if !ok {
		err := accountRepository.AddAccount(nationAccount)
		if err != nil {
			return err
		}
	}

	liquidationAccount := domain.PaymentDTO{Special: true}
	_, ok = accountRepository.GetAccount(liquidationAccount.AccountNumber)
	if !ok {
		err := accountRepository.AddAccount(liquidationAccount)
		if err != nil {
			return err
		}
	}
	return nil
}

func createPaymentSystem(accountRepository *repository.AccountRepository) *service.PaymentSystem {
	paymentService := service.NewPaymentService(accountRepository)
	emissionSpecialAccountService := service.NewEmissionSpecialAccountService(accountRepository)
	liquidationSpecialAccountService := service.NewLiquidationSpecialAccountService(accountRepository)
	accountService := service.NewAccountService(accountRepository)
	paymentSystem := service.NewPaymentSystem(
		emissionSpecialAccountService,
		liquidationSpecialAccountService,
		accountService,
		paymentService,
	)
	return paymentSystem
}

func Run() {
	config := configs.NewConfig()
	connectDB := db.NewDB(config)
	accountRepository := repository.NewAccountRepository(connectDB)
	err := CreateSpecialAccounts(accountRepository)
	paymentSystemService := createPaymentSystem(accountRepository)

	result, err := paymentSystemService.AddToNationAccount(1000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	fmt.Println("Start Payment System")
}
