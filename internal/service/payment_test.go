package service

import (
	"reflect"
	"testPaymentSystem/configs"
	"testPaymentSystem/internal/db"
	"testPaymentSystem/internal/domain"
	"testPaymentSystem/internal/repository"
	"testing"
	"time"
)

func TestPaymentService_getAccounts(t *testing.T) {
	type fields struct {
		repository *repository.AccountRepository
		config     *configs.Config
	}
	type args struct {
		accountNumberFrom string
		accountNumberTo   string
	}
	config := configs.NewConfig()
	repositoryInstance := repository.NewAccountRepository(db.NewDB(config))
	acc1 := domain.Account{
		AccountNumber: "BY04CBDC36029110100040000000",
		Balance:       0,
		Active:        true,
		Currency:      "BYN",
		Limits:        false,
		CreatedAt:     "2021-09-01 00:00:00",
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
	_ = repositoryInstance.AddAccount(acc1)

	acc2 := domain.Account{
		AccountNumber: "BY04CBDC36029110100040000001",
		Balance:       0,
		Active:        true,
		Currency:      "BYN",
		Limits:        false,
		CreatedAt:     "2021-09-01 00:00:00",
		UpdatedAt:     time.Now().Format("2006-01-02 15:04:05"),
	}
	_ = repositoryInstance.AddAccount(acc2)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   domain.Account
		want1  domain.Account
		want2  bool
	}{
		{
			name: "Test getAccounts Success",
			fields: fields{
				repository: repositoryInstance,
				config:     config,
			},
			args: args{
				accountNumberFrom: "BY04CBDC36029110100040000000",
				accountNumberTo:   "BY04CBDC36029110100040000001",
			},
			want:  acc1,
			want1: acc2,
			want2: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaymentService{
				repository: tt.fields.repository,
				config:     tt.fields.config,
			}
			got, got1, got2 := p.getAccounts(tt.args.accountNumberFrom, tt.args.accountNumberTo)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAccounts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getAccounts() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("getAccounts() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestPaymentService_transactionValidation(t *testing.T) {
	type fields struct {
		repository *repository.AccountRepository
		config     *configs.Config
	}
	type args struct {
		accountFrom domain.Account
		accountTo   domain.Account
		sum         float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test Special account",
			fields: fields{
				repository: &repository.AccountRepository{},
				config:     &configs.Config{},
			},
			args: args{
				accountFrom: domain.Account{
					AccountNumber: "BY04CBDC36029110100040000000",
					Balance:       100,
					Active:        true,
					Currency:      "BYN",
					Limits:        false,
					Special:       true,
				},
				accountTo: domain.Account{
					AccountNumber: "BY04CBDC36029110100040000001",
					Balance:       100,
					Active:        true,
					Currency:      "BYN",
					Limits:        false,
					Special:       true,
				},
				sum: 50,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Test Success transaction validation",
			fields: fields{
				repository: &repository.AccountRepository{},
				config:     &configs.Config{},
			},
			args: args{
				accountFrom: domain.Account{
					AccountNumber: "BY04CBDC36029110100040000000",
					Balance:       100,
					Active:        true,
					Currency:      "BYN",
					Limits:        false,
				},
				accountTo: domain.Account{
					AccountNumber: "BY04CBDC36029110100040000001",
					Balance:       100,
					Active:        true,
					Currency:      "BYN",
					Limits:        false,
				},
				sum: 50,
			},
			want:    true,
			wantErr: false,
		},
		{
			// Test InActive account
			name: "Test InActive account",
			fields: fields{
				repository: &repository.AccountRepository{},
				config:     &configs.Config{},
			},
			args: args{
				accountFrom: domain.Account{
					AccountNumber: "BY04CBDC36029110100040000000",
					Balance:       100,
					Active:        false,
					Currency:      "BYN",
					Limits:        false,
				},
				accountTo: domain.Account{
					AccountNumber: "BY04CBDC36029110100040000001",
					Balance:       100,
					Active:        true,
					Currency:      "BYN",
					Limits:        false,
				},
				sum: 50,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Test enough money",
			fields: fields{
				repository: &repository.AccountRepository{},
				config:     &configs.Config{},
			},
			args: args{
				accountFrom: domain.Account{
					AccountNumber: "BY04CBDC36029110100040000000",
					Balance:       0,
					Active:        true,
					Currency:      "BYN",
					Limits:        false,
				},
				accountTo: domain.Account{
					AccountNumber: "BY04CBDC36029110100040000001",
					Balance:       100,
					Active:        true,
					Currency:      "BYN",
					Limits:        false,
				},
				sum: 100,
			},
			want:    false,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaymentService{
				repository: tt.fields.repository,
				config:     tt.fields.config,
			}
			got, err := p.transactionValidation(tt.args.accountFrom, tt.args.accountTo, tt.args.sum)
			if (err != nil) != tt.wantErr {
				t.Errorf("transactionValidation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("transactionValidation() got = %v, want %v", got, tt.want)
			}
		})
	}
}
