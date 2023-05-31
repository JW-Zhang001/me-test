package query

import (
	"fmt"
	"strings"
	"testing"

	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"

	"me-test/config"
)

var (
	DelegateAddr  = config.QueryStateTestData["DelegateAddr"]
	DepositAddr   = config.QueryStateTestData["DepositAddr"]
	KYCAddr       = config.QueryStateTestData["KYCAddr"]
	RegionID      = config.QueryStateTestData["RegionID"]
	ValidatorAddr = config.QueryStateTestData["ValidatorAddr"]
)

func TestGetChainExistRegionID(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"TestGetChainExistRegionID", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChainExistRegionID()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChainExistRegionID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("GetChainExistRegionID got: ", got)
		})
	}
}

func TestGetChainNotExistNodeID(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"TestGetChainNotExistNodeID", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetChainNotExistNodeID()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetChainNotExistNodeID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("GetChainNotExistNodeID got: ", got)
		})
	}
}

func TestIsStringInSlice(t *testing.T) {
	type args struct {
		s     string
		slice []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"TestIsStringInSlice", args{"test", []string{"test", "test1"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStringInSlice(tt.args.s, tt.args.slice); got != tt.want {
				t.Errorf("IsStringInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandNodeID(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"TestRandNodeID", "node"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandNodeID(config.NodeNumber); !strings.Contains(got, tt.want) {
				t.Errorf("RandNodeID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStakeQueryClient_Delegation(t *testing.T) {
	type args struct {
		delAddr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestStakeQueryClient_Delegation", args{DelegateAddr}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &StakeQueryClient{StakeQuery.Cli, StakeQuery.Ctx}
			got, err := q.Delegation(tt.args.delAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delegation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("Delegation got: ", got)
		})
	}
}

func TestStakeQueryClient_DepositByAcc(t *testing.T) {
	type args struct {
		Addr      string
		queryType stakepb.FixedDepositState
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestStakeQueryClient_DepositByAcc", args{DepositAddr, stakepb.FixedDepositState_AllState}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &StakeQueryClient{StakeQuery.Cli, StakeQuery.Ctx}
			got, err := q.DepositByAcc(tt.args.Addr, tt.args.queryType)
			if (err != nil) != tt.wantErr {
				t.Errorf("DepositByAcc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("DepositByAcc got: ", got)
		})
	}
}

func TestStakeQueryClient_KycList(t *testing.T) {

	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"TestStakeQueryClient_KycList", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &StakeQueryClient{StakeQuery.Cli, StakeQuery.Ctx}
			got, err := q.KycList()
			if (err != nil) != tt.wantErr {
				t.Errorf("KycList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("KycList got: ", got)
		})
	}
}

func TestStakeQueryClient_Regions(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"TestStakeQueryClient_Regions", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &StakeQueryClient{StakeQuery.Cli, StakeQuery.Ctx}
			got, err := q.Regions()
			if (err != nil) != tt.wantErr {
				t.Errorf("Regions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("Regions got: ", got)
		})
	}
}

func TestStakeQueryClient_ShowKyc(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestStakeQueryClient_ShowKyc", args{KYCAddr}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &StakeQueryClient{StakeQuery.Cli, StakeQuery.Ctx}
			got, err := q.ShowKyc(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowKyc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("ShowKyc got: ", got)
		})
	}
}

func TestStakeQueryClient_ShowRegion(t *testing.T) {
	type args struct {
		regionID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestStakeQueryClient_ShowRegion", args{RegionID}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &StakeQueryClient{StakeQuery.Cli, StakeQuery.Ctx}
			got, err := q.ShowRegion(tt.args.regionID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowRegion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("ShowRegion got: ", got)
		})
	}
}

func TestStakeQueryClient_ShowValidator(t *testing.T) {
	type args struct {
		operatorAddr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"TestStakeQueryClient_ShowValidator", args{ValidatorAddr}, "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &StakeQueryClient{StakeQuery.Cli, StakeQuery.Ctx}
			got, err := q.ShowValidator(tt.args.operatorAddr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShowValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("ShowValidator got: ", got.Validator)
		})
	}
}

func TestStakeQueryClient_Validators(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{"TestStakeQueryClient_Validators", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &StakeQueryClient{StakeQuery.Cli, StakeQuery.Ctx}
			got, err := q.Validators()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validators() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println("Validators got: ", got.Validators)
		})
	}
}
