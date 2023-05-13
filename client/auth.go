package client

import (
	"context"
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authpb "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (c *CmClient) GetAccountI(ctx context.Context, address string) (acc authpb.AccountI, err error) {
	req := &authpb.QueryAccountRequest{Address: address}
	res, err := c.AuthClient.Account(ctx, req)
	if err != nil {
		return acc, err
	}

	if err = c.cdc.UnpackAny(res.GetAccount(), &acc); err != nil {
		return acc, err
	}

	return acc, nil
}

// ConvertsAccPrivKey Converts a private key of type secp256k1.
func ConvertsAccPrivKey(privKey string) (*secp256k1.PrivKey, error) {
	priBytes, err := hex.DecodeString(privKey)
	if err != nil {
		return nil, err
	}
	return &secp256k1.PrivKey{Key: priBytes}, nil
}

func GetAccAddress(privKey string) (sdk.AccAddress, error) {
	secp256Pk, err := ConvertsAccPrivKey(privKey)
	if err != nil {
		return nil, err
	}
	return secp256Pk.PubKey().Address().Bytes(), nil
}

func GetAccAddrStr(privKey string) (string, error) {
	addr, err := GetAccAddress(privKey)
	if err != nil {
		return "", err
	}
	return addr.String(), nil
}

func GenAccPriKey() string {
	priKey := secp256k1.GenPrivKey()
	return hex.EncodeToString(priKey.Bytes())
}

//func ConvertsValPrivKey(privKey string) (*ed25519.PrivKey, error) {
//	priBytes, err := hex.DecodeString(privKey)
//	if err != nil {
//		return nil, err
//	}
//	return &ed25519.PrivKey{Key: priBytes}, nil
//}
//
//func GetValAddress(privKey string) (sdk.ValAddress, error) {
//	ed25519Pk, err := ConvertsValPrivKey(privKey)
//	if err != nil {
//		return nil, err
//	}
//	return ed25519Pk.PubKey().Address().Bytes(), nil
//}
