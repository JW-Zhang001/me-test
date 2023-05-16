package bank

import (
	"context"
	"me-test/client"
)

type Keeper struct {
	cil *client.CmClient
	ctx context.Context
}
