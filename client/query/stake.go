package query

import (
	"context"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"me-test/client"
	"me-test/config"
)

func NewStakeQuery() (*Query, context.CancelFunc) {
	var ctx, cancel = context.WithTimeout(context.Background(), config.Timeout)

	var c, err = client.NewCmClient("")
	if err != nil {
		return nil, cancel
	}
	return &Query{Cli: c, Ctx: ctx}, cancel
}

func (q *Query) Delegation(ctx context.Context, delAddr string) (*stakepb.QueryDelegationResponse, error) {
	req := &stakepb.QueryDelegationRequest{DelegatorAddr: delAddr, ValidatorAddr: ""}
	rpcRes, err := q.Cli.StakeClient.Delegation(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *Query) DepositByAcc(ctx context.Context, Addr string, queryType stakepb.FixedDepositState) (*stakepb.QueryFixedDepositByAcctResponse, error) {
	req := &stakepb.QueryFixedDepositByAcctRequest{Account: Addr, QueryType: queryType}
	rpcRes, err := q.Cli.StakeClient.FixedDepositByAcct(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *Query) KycList(ctx context.Context) (*stakepb.QueryAllKycResponse, error) {
	req := &stakepb.QueryAllKycRequest{}
	rpcRes, err := q.Cli.StakeClient.KycAll(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}
