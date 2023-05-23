package query

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

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

func (q *Query) ShowKyc(ctx context.Context, addr string) (*stakepb.QueryGetKycResponse, error) {
	req := &stakepb.QueryGetKycRequest{Account: addr}
	rpcRes, err := q.Cli.StakeClient.Kyc(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *Query) ShowRegion(ctx context.Context, regionID string) (*stakepb.QueryGetRegionResponse, error) {
	req := &stakepb.QueryGetRegionRequest{RegionId: regionID}
	rpcRes, err := q.Cli.StakeClient.Region(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *Query) ShowValidator(ctx context.Context, operatorAddr string) (*stakepb.QueryValidatorResponse, error) {
	req := &stakepb.QueryValidatorRequest{ValidatorAddr: operatorAddr}
	rpcRes, err := q.Cli.StakeClient.Validator(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *Query) Validators(ctx context.Context) (*stakepb.QueryValidatorsResponse, error) {
	req := &stakepb.QueryValidatorsRequest{}
	rpcRes, err := q.Cli.StakeClient.Validators(ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func RandNodeID() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	num := r.Intn(16) + 1
	nodeID := "node" + strconv.Itoa(num)

	return nodeID
}

func GetChainNotExistNodeID() (string, error) {
	val, err := StakeQuery.Validators(StakeQuery.Ctx)
	if err != nil {
		return "", err
	}

	validators := make([]string, 0, 100)
	for _, v := range val.Validators {
		validators = append(validators, v.Description.Moniker)
	}

	timeout := time.After(10 * time.Second)      // Set a timeout period of 10 seconds
	ticker := time.Tick(1000 * time.Millisecond) // It's triggered every 1000 milliseconds

	for {
		select {
		case <-timeout:
			return "", errors.New("timeout occurred")
		case <-ticker:
			nodeID := RandNodeID()
			ok := IsStringInSlice(nodeID, validators)
			if !ok {
				fmt.Println("Chain not exist nodeID:", nodeID)
				return nodeID, nil
			}
		}
	}
}

func IsStringInSlice(s string, slice []string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}
