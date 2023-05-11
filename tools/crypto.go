package tools

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetPrivKey 根据导入私钥字符串转换为 secp256k1.PrivKey 类型的私钥
func GetPrivKey(privKey string) (*secp256k1.PrivKey, error) {
	priBytes, err := hex.DecodeString(privKey)
	if err != nil {
		return nil, err
	}
	return &secp256k1.PrivKey{Key: priBytes}, nil
}

func GetAccAddr(privKey string) (sdk.AccAddress, error) {
	pk, err := GetPrivKey(privKey)
	if err != nil {
		return nil, err
	}
	return pk.PubKey().Address().Bytes(), nil
}

func GenPriKey() string {
	priKey := secp256k1.GenPrivKey()
	return hex.EncodeToString(priKey.Bytes())
}
