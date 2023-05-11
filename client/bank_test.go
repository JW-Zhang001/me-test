package client

import (
	"context"
	"math/big"
	"strings"
	"testing"
)

var (
	ctx  = context.Background()
	c, _ = NewCmClient("")
)

func TestCmClient_Balance(t *testing.T) {

	type args struct {
		ctx  context.Context
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"case1", args{ctx, "cosmos1fap8hp3t3xt20qw4sczlyrk6n92uffj4r4kw77"}, "umec", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.Balance(tt.args.ctx, tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Balance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotStr := got.Balance.String()
			if strings.Contains(gotStr, tt.want) {
				t.Log(gotStr)
			} else {
				t.Errorf("Balance() not contains got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCmClient_AllBalances(t *testing.T) {

	type args struct {
		ctx  context.Context
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"case1", args{ctx, "cosmos1gmpxkchcdgfq995zye5efwzfw86zfa4vt4ke8g"}, "umec", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := c.AllBalances(tt.args.ctx, tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("AllBalances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotStr := got.Balances.String()
			if strings.Contains(gotStr, tt.want) {
				t.Log(gotStr)
			} else {
				t.Errorf("Balance() not contains got = %v, want %v", got, tt.want)
			}
		})
	}
}

// total supply of all coins.
func TestCmClient_TotalSupply(t *testing.T) {
	a := big.NewInt(10000000000)
	b := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(6), nil)
	stakeCoin := big.NewInt(0).Mul(a, b)

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{name: "case1", args: args{ctx}, want: stakeCoin, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.TotalSupply(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("TotalSupply() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotUint64 := got.Supply[0].Amount.Uint64()
			wantUint64 := tt.want.Uint64()
			if gotUint64 >= wantUint64 {
				t.Logf("Total Supply: %v", gotUint64)
			} else {
				t.Errorf("Total Supply got = %v, want >= %v", gotUint64, wantUint64)
			}
		})
	}
}
