package query

import (
	"context"
	"fmt"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"testing"
)

func TestCmClient_Delegation(t *testing.T) {

	type args struct {
		ctx     context.Context
		delAddr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestCmClient_Delegation", args{StakeQuery.Ctx, "cosmos1cg77vlldvxr3se38quuzrpu5guum9l0ct73r6u"}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StakeQuery.Delegation(tt.args.ctx, tt.args.delAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delegation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestCmClient_DepositByAcc(t *testing.T) {

	type args struct {
		ctx       context.Context
		addr      string
		queryType stakepb.FixedDepositState
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestCmClient_DepositByAcc", args{StakeQuery.Ctx, "cosmos1yp07x6wjruw28066p9nfugzdrvxyxgu09frxkx", stakepb.FixedDepositState_Expired}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StakeQuery.DepositByAcc(tt.args.ctx, tt.args.addr, tt.args.queryType)
			if (err != nil) != tt.wantErr {
				t.Errorf("DepositByAcc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got.FixedDeposit[0].Id)
		})
	}
}

func TestQuery_Validators(t *testing.T) {

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string

		args    args
		want    string
		wantErr bool
	}{
		{"case3", args{StakeQuery.Ctx}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StakeQuery.Validators(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validators() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got.Validators)
		})
	}
}
