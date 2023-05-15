package client

import (
	"context"
	"fmt"
	authpb "github.com/cosmos/cosmos-sdk/x/auth/types"
	"testing"
)

func TestCmClient_GetAccountI(t *testing.T) {

	type args struct {
		ctx     context.Context
		address string
	}
	tests := []struct {
		name    string
		args    args
		wantAcc authpb.AccountI
		wantErr bool
	}{
		{"case1", args{ctx, "cosmos1k0nfwtzsv30xtdxturftyga9ajjsnhq9vh3kya"},
			&authpb.BaseAccount{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotAcc, err := c.GetAccountI(tt.args.ctx, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotAcc.GetAddress().String())
		})
	}
}
