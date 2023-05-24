package client

import (
	"context"
	"go.uber.org/zap"
	"sync/atomic"
)

type AccountInfo struct {
	Sequence  uint64
	AccNumber uint64
}

func NewAccountInfo(ctx context.Context, c *CmClient, addr string) (AccountInfo, error) {
	sequence, err := GetSequence(ctx, c, addr)
	if err != nil {
		zap.S().Error("GetSequence is err:", err)
		return AccountInfo{}, err
	}
	accNumber, err := GetAccNumber(ctx, c, addr)
	if err != nil {
		zap.S().Error("GetAccNumber is err:", err)
		return AccountInfo{}, err
	}

	return AccountInfo{sequence, accNumber}, nil
}

func GetSequence(ctx context.Context, c *CmClient, address string) (uint64, error) {
	acc, err := c.GetAccountI(ctx, address)
	if err != nil {
		zap.S().Error("GetSequence is err:", err)
		return 0, err
	}
	return acc.GetSequence(), nil
}

func GetAccNumber(ctx context.Context, c *CmClient, address string) (uint64, error) {
	acc, err := c.GetAccountI(ctx, address)
	if err != nil {
		zap.S().Error("GetAccNumber is err:", err)
		return 0, err
	}
	return acc.GetAccountNumber(), nil
}

func (a *AccountInfo) IncrementSequence() uint64 {
	return atomic.AddUint64(&a.Sequence, 1)
}

func (a *AccountInfo) ResetSequence(nonce uint64) {
	atomic.StoreUint64(&a.Sequence, nonce)
}

func (a *AccountInfo) CurrentSequence() uint64 {
	return atomic.LoadUint64(&a.Sequence)
}

type Wallet struct {
	ID      uint64
	PrivKey string
	Addr    string
	AccInfo AccountInfo
}

func NewWallet(ctx context.Context, c *CmClient) (Wallet, error) {
	acc, err := GenWalletAcc()
	if err != nil {
		return Wallet{}, err
	}
	accInfo, _ := NewAccountInfo(ctx, c, acc["Addr"])
	return Wallet{0, acc["PrivKey"], acc["Addr"], accInfo}, nil
}

func ImportWallet(ctx context.Context, c *CmClient, privKey string, id uint64) (Wallet, error) {
	addr, _ := GetAccAddrStr(privKey)
	accInfo, _ := NewAccountInfo(ctx, c, addr)
	return Wallet{id, privKey, addr, accInfo}, nil
}
