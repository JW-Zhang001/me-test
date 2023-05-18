package staking

import (
	"context"
	"me-test/client"
	"me-test/config"
)

type Keeper struct {
	Cli *client.CmClient
	Ctx context.Context
}

func NewKeeper() (*Keeper, context.CancelFunc) {
	var ctx, cancel = context.WithTimeout(context.Background(), config.Timeout)

	var c, err = client.NewCmClient("")
	if err != nil {
		return nil, cancel
	}
	return &Keeper{Cli: c, Ctx: ctx}, cancel
}
