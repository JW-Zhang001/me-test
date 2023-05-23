package staking

import (
	"fmt"
	"me-test/config"
	"testing"
)

func TestKeeper_NewValidator(t *testing.T) {
	tmPubK := "{\"type\": \"tendermint/PubKeyEd25519\",\"value\": \"GxyY54js2kyciCnYmhMmHwAJ/dcf56wDB8J8vuShid8=\"}"
	type args struct {
		privKey     string
		tmPubKeyStr string
		coinStr     string
		moniker     string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestKeeper_NewValidator",
			args{config.SuperAdminPrivKey, tmPubK, "100mec", "node5"},
			"", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keeper{Cli: C, Ctx: Ctx}
			got, err := k.NewValidator(tt.args.privKey, tt.args.tmPubKeyStr, tt.args.coinStr, tt.args.moniker)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)

		})
	}
}

func TestKeeper_EditValidator(t *testing.T) {

	type args struct {
		privKey         string
		operatorAddress string
		ownerAddress    string
		moniker         string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"case1", args{config.SuperAdminPrivKey, "cosmosvaloper14y432ss720m8hzpdx4dknet5epv5uc7hqypqwa",
			"cosmos1tpry5qwctzumtyg59guvfgwdjxr8my424vf5qr", "node2"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keeper{Cli: C, Ctx: Ctx}
			got, err := k.EditValidator(tt.args.privKey, tt.args.operatorAddress, tt.args.ownerAddress, tt.args.moniker)
			if (err != nil) != tt.wantErr {
				t.Errorf("EditValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
