package staking

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"me-test/config"
	"testing"
)

func TestKeeper_Delegate(t *testing.T) {

	type args struct {
		delPriKey string
		amount    sdk.Coin
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "TestKeeper_Delegate", args: args{
			"fd7bf03c3fd0e4b70de7c3bde6d8f6267d6bb6f6d5719d24003adb0ef30638ea",
			sdk.NewInt64Coin(sdk.BaseMEDenom, config.DelegateAmount),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keeper{Cli: C, Ctx: Ctx}
			got, err := k.Delegate(tt.args.delPriKey, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delegate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
