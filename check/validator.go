package check

import (
	"fmt"
	"strconv"

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
	stakeTokensPool, err := q.BankQuery.Balance(config.ModuleAccountList["stake_tokens_pool"])
	if err != nil {
		return nil, err
	}
	bondedStakeTokensPool, err := q.BankQuery.Balance(config.ModuleAccountList["bonded_stake_tokens_pool"])
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
		stakeAmt, err := strconv.ParseInt(config.ValidatorStakeAmount[:len(config.ValidatorStakeAmount)-3], 10, 64)
		if err != nil {
			zap.S().Error("Conversion failed", err)
			return "", err
		}
		uStakeAmt := stakeAmt * 1000000
		if actual1 != uStakeAmt {
			return "Assert false", fmt.Errorf("CheckerNewValidator assert false error = %v, wantErr %v", actual1, uStakeAmt)
		}
		actual2 := laterBalancesList[1].Balance.Amount.Int64() - balancesList[1].Balance.Amount.Int64()
		if actual2 != uStakeAmt {
			return "Assert false", fmt.Errorf("CheckerNewValidator assert false error = %v, wantErr %v", actual2, uStakeAmt)
		}
		return result, err
	}
}
