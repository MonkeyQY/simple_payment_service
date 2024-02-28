package service

import (
	"reflect"
	"testPaymentSystem/configs"
	"testPaymentSystem/internal/domain"
	"testPaymentSystem/internal/repository"
	"testing"
)

func TestPaymentService_Send(t *testing.T) {
	type fields struct {
		repository *repository.AccountRepository
		config     *configs.Config
	}
	type args struct {
		accountNumberFrom string
		accountNumberTo   string
		sum               float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PaymentService{
				repository: tt.fields.repository,
				config:     tt.fields.config,
			}
			got, err := p.Send(tt.args.accountNumberFrom, tt.args.accountNumberTo, tt.args.sum)
			if (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Send() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPaymentService_getAccounts(t *testing.T) {
	type fields struct {
		repository *repository.AccountRepository
		config     *configs.Config
	}
	type args struct {
		accountNumberFrom string
		accountNumberTo   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   domain.PaymentDTO
		want1  domain.PaymentDTO
		want2  bool
	}{
		// TODO: Add test cases.
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
		accountFrom domain.PaymentDTO
		accountTo   domain.PaymentDTO
		sum         float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
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
