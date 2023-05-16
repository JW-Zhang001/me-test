package staking

import (
	"context"
	"me-test/client"
	"me-test/config"
)

type Keeper struct {
	Cil *client.CmClient
	Ctx context.Context
}

func NewKeeper() (*Keeper, context.CancelFunc) {
	var Ctx, cancel = context.WithTimeout(context.Background(), config.Timeout)

	// defer cancel()
	var C, err = client.NewCmClient("")
	if err != nil {
		return nil, cancel
	}
	return &Keeper{Cil: C, Ctx: Ctx}, cancel
}
