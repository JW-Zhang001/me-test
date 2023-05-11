package client

import (
	"context"
	authpb "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (c *CmClient) GetAccountI(ctx context.Context, address string) (acc authpb.AccountI, err error) {
	req := &authpb.QueryAccountRequest{Address: address}
	res, err := c.AuthClient.Account(ctx, req)
	if err != nil {
		return acc, err
	}

	if err = c.cdc.UnpackAny(res.GetAccount(), &acc); err != nil {
		return acc, err
	}

	return
}
