package query

import (
	"context"
)

var (
	BankQuery  *Query
	StakeQuery *Query

	Cancel context.CancelFunc
)

func init() {
	BankQuery, Cancel = NewBankQuery()
	StakeQuery, _ = NewBankQuery()
}
