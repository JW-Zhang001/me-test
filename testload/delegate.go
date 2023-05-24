package testload

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.uber.org/zap"
	"me-test/client"
	"me-test/client/query"
	"me-test/config"
	"me-test/testcase"
)

var (
	k = testcase.StakeKeeper
)

func Delegate(privKey string, amount int64, sequence uint64) (*txpb.BroadcastTxResponse, error) {
	fromAccAddr, _ := client.GetAccAddress(privKey)

	// valAddr and valStr Pass the null value to determine the kyc status of the user based on delAccAddr
	msg := stakepb.NewMsgDelegate(fromAccAddr, sdk.ValAddress{}, sdk.NewInt64Coin(sdk.DefaultBondDenom, amount), "")
	if msg.ValidateBasic() != nil {
		return nil, fmt.Errorf("msg.ValidateBasic error: %v", msg.ValidateBasic())
	}

	res, err := k.Cli.SendBroadcastCheckTx(k.Ctx, privKey, msg, sequence, config.DefaultFees)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func DelegateBenchmark() {
	wallets, err := initWallet()
	if err != nil {
		zap.S().Error("InitWallet error: ", err)
	}
	for _, wal := range wallets {
		if wal.ID == 1 {
			sequence := wal.AccInfo.Sequence
			res1, err := Delegate(wal.PrivKey, config.DelegateAmount, sequence)
			if err != nil {
				zap.S().Error("Delegate error: ", err)
			}
			zap.S().Info("----res1: ", res1)

			sequence = wal.AccInfo.IncrementSequence()
			res2, err := Delegate(wal.PrivKey, config.DelegateAmount, sequence)
			if err != nil {
				zap.S().Error("Delegate error: ", err)
			}
			zap.S().Info("----res2: ", res2)

			sequence = wal.AccInfo.IncrementSequence()
			res3, err := Delegate(wal.PrivKey, config.DelegateAmount, sequence)
			if err != nil {
				zap.S().Error("Delegate error: ", err)
			}
			zap.S().Info("----res3: ", res3)
		}
	}
}

func initWallet() (wallets []client.Wallet, err error) {
	superAdminWal, _ := client.ImportWallet(k.Ctx, k.Cli, config.SuperAdminPrivKey, 0)
	wallets = append(wallets, superAdminWal)

	user1Acc, _ := client.GenWalletAcc()

	if err = testcase.TestTx(config.SuperAdminPrivKey, user1Acc["Addr"], config.TxAmount); err != nil {
		return wallets, err
	}

	user1Wal, _ := client.ImportWallet(k.Ctx, k.Cli, user1Acc["PrivKey"], 1)
	wallets = append(wallets, user1Wal)

	zap.S().Info("InitWallet success:", wallets)
	return wallets, nil
}

func MixDelegateBenchmark() {
	wallets, err := initWallet()
	if err != nil {
		zap.S().Error("InitWallet error: ", err)
	}
	var superadminWal client.Wallet
	for _, wal := range wallets {
		if wal.ID == 0 {
			superadminWal = wal
		}
		if wal.ID == 1 {
			sequence := wal.AccInfo.Sequence
			res1, err := Delegate(wal.PrivKey, config.DelegateAmount, sequence)
			if err != nil {
				zap.S().Error("Delegate error: ", err)
			}
			zap.S().Info("----Not Kyc Delegate res1: ", res1)

			regionID, err := query.GetChainExistRegionID()
			if err != nil {
				zap.S().Error("GetChainExistRegionID error: ", err)
				return
			}
			// InitWallet had a transfer, so the sequence needs to increment
			sequence = superadminWal.AccInfo.IncrementSequence()
			res2, err := NewKyc(superadminWal.PrivKey, wal.Addr, regionID, sequence)
			if err != nil {
				zap.S().Error("NewKyc error: ", err)
			}
			zap.S().Infof("----Superadmin New Kyc to %v, res2: %v", wal.Addr, res2)

			sequence = wal.AccInfo.IncrementSequence()
			res3, err := Delegate(wal.PrivKey, config.DelegateAmount, sequence)
			if err != nil {
				zap.S().Error("Kyc Delegate error: ", err)
			}
			zap.S().Info("----Kyc Delegate res3: ", res3)
		}
	}
}

func MixDelegateBenchmark2() {
	wallets, err := initWallet()
	if err != nil {
		zap.S().Error("InitWallet error: ", err)
	}
	var superadminWal client.Wallet
	for _, wal := range wallets {
		if wal.ID == 0 {
			superadminWal = wal
		}
		if wal.ID == 1 {
			sequence := wal.AccInfo.Sequence
			res1, err := Delegate(wal.PrivKey, config.DelegateAmount, sequence)
			if err != nil {
				zap.S().Error("Delegate error: ", err)
			}
			zap.S().Info("----Not Kyc Delegate res1: ", res1)

			sequence = wal.AccInfo.IncrementSequence()
			res2, err := Delegate(wal.PrivKey, config.DelegateAmount, sequence)
			if err != nil {
				zap.S().Error("Delegate error: ", err)
			}
			zap.S().Info("----Not Kyc Delegate res2: ", res2)

			regionID, err := query.GetChainExistRegionID()
			if err != nil {
				zap.S().Error("GetChainExistRegionID error: ", err)
				return
			}
			// InitWallet had a transfer, so the sequence needs to increment
			sequence = superadminWal.AccInfo.IncrementSequence()
			res3, err := NewKyc(superadminWal.PrivKey, wal.Addr, regionID, sequence)
			if err != nil {
				zap.S().Error("NewKyc error: ", err)
			}
			zap.S().Infof("----Superadmin New Kyc to %v, res3: %v", wal.Addr, res3)
		}
	}
}
