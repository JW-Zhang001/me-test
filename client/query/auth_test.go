package query

import (
	"fmt"
	"me-test/config"
	"testing"
)

func TestAuthQueryClient_Account(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestAuthQueryClient_Account", args{config.BaseAccountList["PM"]}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &AuthQueryClient{BankQuery.Cli, BankQuery.Ctx}
			got, err := q.Account(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Account() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got.Account)
		})
	}
}

func TestAuthQueryClient_Accounts(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"TestAuthQueryClient_Accounts", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &AuthQueryClient{BankQuery.Cli, BankQuery.Ctx}
			got, err := q.Accounts()
			if (err != nil) != tt.wantErr {
				t.Errorf("Accounts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got.Accounts)
		})
	}
}
