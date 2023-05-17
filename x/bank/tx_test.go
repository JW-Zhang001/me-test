package bank

import (
	"context"
	"fmt"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"me-test/client"
	"me-test/config"
	"testing"
)

var ctx, cancel = context.WithTimeout(context.Background(), config.Timeout)

// defer cancel()
var c, _ = client.NewCmClient("")

func TestKeeper_TransferTx(t *testing.T) {

	type args struct {
		privKey string
		toAddr  string
		amount  types.Coin
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestKeeper_TransferTx",
			args{"1dd9136a57ae1825234d5e820a1e25b9d292e82f9c75255b4426bdf6827efce2",
				"cosmos18cg9awlkpy2upsq380a8vhwa0cnppn68nsmqxp",
				sdk.NewInt64Coin(sdk.BaseMEDenom, 1000000)},
			"", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Keeper{Cil: c, Ctx: ctx}
			got, err := s.TransferTx(tt.args.privKey, tt.args.toAddr, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransferTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestKeeper_TxToAdmin(t *testing.T) {

	type args struct {
		privKey string
		amount  sdk.Coin
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestKeeper_TxToAdmin", args{config.SuperAdminPrivKey, sdk.NewInt64Coin(sdk.BaseMEDenom, 100000000)}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Keeper{Cil: c, Ctx: ctx}

			got, err := s.TxToAdmin(tt.args.privKey, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("TxToAdmin() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
