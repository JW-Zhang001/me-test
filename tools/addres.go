package tools

import (
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/rand"
	"strconv"
	"time"
)

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

func GenAccPrivKey() string {
	priKey := secp256k1.GenPrivKey()
	return hex.EncodeToString(priKey.Bytes())
}

func GenWalletAcc() (map[string]string, error) {
	walletAcc := make(map[string]string, 5)
	privKey := GenAccPrivKey()
	addr, _ := GetAccAddrStr(privKey)

	walletAcc["PrivKey"] = privKey
	walletAcc["Addr"] = addr
	return walletAcc, nil
}

func RandNodeID() string {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	num := r.Intn(20) + 1
	nodeID := "node" + strconv.Itoa(num)

	return nodeID
}
