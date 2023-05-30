package query

import (
	"math"
	"me-test/config"
	"reflect"
	"testing"
)

func TestBankQueryClient_AllBalances(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestQueryAllBalances", args{addr: config.ModuleAccountList["treasury_pool"]}, "umec", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &BankQueryClient{BankQuery.Cli, BankQuery.Ctx}
			got, err := q.AllBalances(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllBalances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			denom := got.Balances[0].Denom
			if denom != tt.want {
				t.Errorf("AllBalances() got = %v, want %v", denom, tt.want)
			}
		})
	}
}

func TestBankQueryClient_Balance(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestQueryBalance", args{addr: config.ModuleAccountList["treasury_pool"]}, "umec", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &BankQueryClient{BankQuery.Cli, BankQuery.Ctx}
			got, err := q.Balance(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Balance.Denom, tt.want) {
				t.Errorf("Balance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// total supply of all coins.
func TestBankQueryClient_TotalSupply(t *testing.T) {

	tests := []struct {
		id      uint8
		name    string
		want    uint64
		wantErr bool
	}{
		{1, "TestQueryTotalSupply", uint64(100 * math.Pow10(8) * math.Pow10(6)), false},
		{2, "TestQueryTotalSupply", uint64(200 * math.Pow10(8) * math.Pow10(6)), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &BankQueryClient{BankQuery.Cli, BankQuery.Ctx}
			got, err := q.TotalSupply()
			if (err != nil) != tt.wantErr {
				t.Errorf("TotalSupply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.id == 1 {
				if got.Supply[0].Amount.Uint64() <= tt.want {
					t.Errorf("TotalSupply() got = %v, want %v", got, tt.want)
					return
				}
			}
			if tt.id == 2 {
				if got.Supply[0].Amount.Uint64() > tt.want {
					t.Errorf("TotalSupply() got = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
