package check

import (
	"fmt"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	"go.uber.org/zap"
	q "me-test/client/query"
	"me-test/config"
)

/*
QueryStakeTokensPool
@Description: get stakeTokensPool and bondedStakeTokensPool balance
@return balancesList
@return err
*/
func QueryStakeTokensPool() (balancesList []*bankpb.QueryBalanceResponse, err error) {
	stakeTokensPool, err := q.BankQuery.Balance(q.BankQuery.Ctx, config.ModuleAccountList["stake_tokens_pool"])
	if err != nil {
		return nil, err
	}
	bondedStakeTokensPool, err := q.BankQuery.Balance(q.BankQuery.Ctx, config.ModuleAccountList["bonded_stake_tokens_pool"])
	if err != nil {
		return nil, err
	}
	balancesList = append(balancesList, stakeTokensPool, bondedStakeTokensPool)
	return balancesList, nil
}

/*
CheckerNewValidator
@Description: Use decorator to realize before and after balance amount verification
@param fn
@return func(nodeID, coinStr string) (string, error)
*/
func CheckerNewValidator(fn func(nodeID, coinStr string) (string, error)) func(nodeID, coinStr string) (string, error) {
	return func(nodeID, coinStr string) (string, error) {
		balancesList, err := QueryStakeTokensPool()
		if err != nil {
			zap.S().Errorf("NewValidator before get balance error: %v", err)
			return "", err
		}
		zap.S().Info("NewValidator before: ", balancesList)

		result, err := fn(nodeID, coinStr)

		laterBalancesList, err := QueryStakeTokensPool()
		if err != nil {
			zap.S().Errorf("NewValidator later get balance error: %v", err)
			return "", err
		}
		zap.S().Info(laterBalancesList)

		actual1 := balancesList[0].Balance.Amount.Int64() - laterBalancesList[0].Balance.Amount.Int64()
		if actual1 != int64(1000000000) {
			return "Assert false", fmt.Errorf("assert false error = %v, wantErr %v", actual1, int64(1000000000))
		}
		actual2 := laterBalancesList[1].Balance.Amount.Int64() - balancesList[1].Balance.Amount.Int64()
		if actual2 != int64(1000000000) {
			return "Assert false", fmt.Errorf("assert false error = %v, wantErr %v", actual2, int64(1000000000))
		}
		return result, err
	}
}
