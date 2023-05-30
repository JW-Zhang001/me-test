package query

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"time"

	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"

	"me-test/client"
	"me-test/config"
)

type StakeQueryClient struct {
	Cli *client.CmClient
	Ctx context.Context
}

func NewStakeQuery() (*StakeQueryClient, context.CancelFunc) {
	var ctx, cancel = context.WithTimeout(context.Background(), config.Timeout)

	var c, err = client.NewCmClient("")
	if err != nil {
		return nil, cancel
	}
	return &StakeQueryClient{Cli: c, Ctx: ctx}, cancel
}

func (q *StakeQueryClient) Delegation(delAddr string) (*stakepb.QueryDelegationResponse, error) {
	req := &stakepb.QueryDelegationRequest{DelegatorAddr: delAddr}
	rpcRes, err := q.Cli.StakeClient.Delegation(StakeQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *StakeQueryClient) DepositByAcc(Addr string, queryType stakepb.FixedDepositState) (*stakepb.QueryFixedDepositByAcctResponse, error) {
	req := &stakepb.QueryFixedDepositByAcctRequest{Account: Addr, QueryType: queryType}
	rpcRes, err := q.Cli.StakeClient.FixedDepositByAcct(StakeQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *StakeQueryClient) KycList() (*stakepb.QueryAllKycResponse, error) {
	req := &stakepb.QueryAllKycRequest{}
	rpcRes, err := q.Cli.StakeClient.KycAll(StakeQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *StakeQueryClient) ShowKyc(addr string) (*stakepb.QueryGetKycResponse, error) {
	req := &stakepb.QueryGetKycRequest{Account: addr}
	rpcRes, err := q.Cli.StakeClient.Kyc(StakeQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *StakeQueryClient) ShowRegion(regionID string) (*stakepb.QueryGetRegionResponse, error) {
	req := &stakepb.QueryGetRegionRequest{RegionId: regionID}
	rpcRes, err := q.Cli.StakeClient.Region(StakeQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *StakeQueryClient) Regions() (*stakepb.QueryAllRegionResponse, error) {
	req := &stakepb.QueryAllRegionRequest{}
	rpcRes, err := q.Cli.StakeClient.RegionAll(StakeQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *StakeQueryClient) ShowValidator(operatorAddr string) (*stakepb.QueryValidatorResponse, error) {
	req := &stakepb.QueryValidatorRequest{ValidatorAddr: operatorAddr}
	rpcRes, err := q.Cli.StakeClient.Validator(StakeQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *StakeQueryClient) Validators() (*stakepb.QueryValidatorsResponse, error) {
	req := &stakepb.QueryValidatorsRequest{}
	rpcRes, err := q.Cli.StakeClient.Validators(StakeQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func RandNodeID(number int) string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	num := r.Intn(number) + 1
	nodeID := "node" + strconv.Itoa(num)

	return nodeID
}

func GetChainNotExistNodeID() (string, error) {
	val, err := StakeQuery.Validators()
	if err != nil {
		return "", err
	}

	validators := make([]string, 0, 100)
	for _, v := range val.Validators {
		validators = append(validators, v.Description.Moniker)
	}

	timeout := time.After(60 * time.Second)      // Set a timeout period of 60 seconds
	ticker := time.Tick(1000 * time.Millisecond) // It's triggered every 1000 milliseconds

	for {
		select {
		case <-timeout:
			return "", errors.New("timeout occurred")
		case <-ticker:
			nodeID := RandNodeID(config.NodeNumber)
			ok := IsStringInSlice(nodeID, validators)
			if !ok {
				zap.S().Info("Chain not exist nodeID:", nodeID)
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

func GetChainExistRegionID() (string, error) {
	region, err := StakeQuery.Regions()
	if err != nil {
		zap.S().Errorf("GetChainExistRegionID error: %v", err)
		return "", err
	}
	regionList := region.Region
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	regionInfo := regionList[r.Intn(len(regionList))]
	return regionInfo.RegionId, nil
}
