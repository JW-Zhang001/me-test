package staking

import (
	"fmt"
	"me-test/config"
	_ "me-test/initialize"
	"me-test/tools"
	"testing"
)

func TestKeeper_NewKyc(t *testing.T) {
	toAddr, _ := tools.GetAccAddress(tools.GenAccPriKey())

	type args struct {
		privKey  string
		toAddr   string
		regionId string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestKeeper_NewKyc", args{
			privKey: config.SuperAdminPrivKey, toAddr: toAddr.String(), regionId: "2",
		}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keeper{Cli: C, Ctx: Ctx}
			got, err := k.NewKyc(tt.args.privKey, tt.args.toAddr, tt.args.regionId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewKyc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
