package staking

import (
	"context"
	"me-test/client"
	"me-test/config"
)

type Keeper struct {
	cil *client.CmClient
	ctx context.Context
}

var Ctx, cancel = context.WithTimeout(context.Background(), config.Timeout)

// defer cancel()
var C, _ = client.NewCmClient("")
