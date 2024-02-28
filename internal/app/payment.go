package app

import (
	"fmt"
	"testPaymentSystem/configs"
	"testPaymentSystem/internal/db"
	"testPaymentSystem/internal/domain"
	"testPaymentSystem/internal/repository"
	"testPaymentSystem/internal/service"
	"time"
)

func CreateSpecialAccounts(accountRepository *repository.AccountRepository, config *configs.Config) error {
	nationAccount := domain.PaymentDTO{
		Special:       true,
		AccountNumber: config.NationalAccountNumber,
		Currency:      "BYN",
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
	_, ok := accountRepository.GetAccount(nationAccount.AccountNumber)
	if !ok {
		err := accountRepository.AddAccount(nationAccount)
		if err != nil {
			return err
		}
	}

	liquidationAccount := domain.PaymentDTO{
		Special:       true,
		AccountNumber: config.LiquidationAccountNumber,
		Currency:      "BYN",
		CreatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
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
	err := CreateSpecialAccounts(accountRepository, config)
	paymentSystemService := createPaymentSystem(accountRepository)

	result, err := paymentSystemService.AddToNationAccount(1000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	result, err = paymentSystemService.AddToLiquidationAccount(1000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	fmt.Println(paymentSystemService.GetNationAccountNumber())
	fmt.Println(paymentSystemService.GetLiquidationAccountNumber())

	acc1, err := paymentSystemService.CreateNewAccount()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(acc1)

	acc2, err := paymentSystemService.CreateNewAccount()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(acc2)

	accounts := paymentSystemService.GetAllAccounts()
	for _, account := range accounts {
		fmt.Printf("%+v \n", account)
	}

	_, err = paymentSystemService.Transfer(acc1, acc2, 100)
	if err != nil {
		fmt.Println(err)
	}

	balance, err := paymentSystemService.Replenishment(acc1, 1000)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Balance acc1 after replenishment", balance)

	isSuccess, err := paymentSystemService.Transfer(acc1, acc2, 100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Transfer result:", isSuccess)

	isSuccess, err = paymentSystemService.Transfer(acc1, "dsfg", 100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Transfer result:", isSuccess)

	isSuccess, err = paymentSystemService.Transfer("dsfg", acc2, 100)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Transfer result:", isSuccess)

}
