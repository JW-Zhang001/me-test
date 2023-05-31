package staking

import (
	"fmt"
	"me-test/config"
	"testing"
)

func TestKeeper_NewRegion(t *testing.T) {

	type args struct {
		privKey   string
		regionId  string
		name      string
		validator string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestKeeper_NewRegion", args{regionId: "1", name: "TCA",
			validator: "cosmosvaloper1q0hfp4364h4gxcantag3qgam3t04pldku67dm2", privKey: config.SuperAdminPrivKey},
			"", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k := &Keeper{Cli: C, Ctx: Ctx}
			got, err := k.NewRegion(tt.args.privKey, tt.args.regionId, tt.args.name, tt.args.validator)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRegion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
