package query

import (
	"context"
	"fmt"

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
