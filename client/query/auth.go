package query

import (
	"context"
	authpb "github.com/cosmos/cosmos-sdk/x/auth/types"

	"me-test/client"
	"me-test/config"
)

var (
	BankQuery  *BankQueryClient
	StakeQuery *StakeQueryClient
	AuthQuery  *AuthQueryClient
)

func init() {
	BankQuery, _ = NewBankQuery()
	StakeQuery, _ = NewStakeQuery()
	AuthQuery, _ = NewAuthQuery()
}

type AuthQueryClient struct {
	Cli *client.CmClient
	Ctx context.Context
}

func NewAuthQuery() (*AuthQueryClient, context.CancelFunc) {
	var ctx, cancel = context.WithTimeout(context.Background(), config.Timeout)

	var c, err = client.NewCmClient("")
	if err != nil {
		return nil, cancel
	}
	return &AuthQueryClient{Cli: c, Ctx: ctx}, cancel
}

func (q *AuthQueryClient) Accounts() (*authpb.QueryAccountsResponse, error) {
	req := &authpb.QueryAccountsRequest{}
	rpcRes, err := q.Cli.AuthClient.Accounts(AuthQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}

func (q *AuthQueryClient) Account(addr string) (*authpb.QueryAccountResponse, error) {
	req := &authpb.QueryAccountRequest{Address: addr}
	rpcRes, err := q.Cli.AuthClient.Account(AuthQuery.Ctx, req)
	if err != nil {
		return nil, err
	}
	return rpcRes, nil
}
