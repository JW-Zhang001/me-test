package bank

import (
	"context"
	"me-test/client"
)

type Keeper struct {
	Cil *client.CmClient
	Ctx context.Context
}
