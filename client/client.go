package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/tendermint/tendermint/rpc/client/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"me-test/config"
)

type CmClient struct {
	Conn    *grpc.ClientConn
	CliHTTP *http.HTTP

	cdc      *codec.ProtoCodec
	txConfig client.TxConfig
}

/*
NewCmClient
rpcURI 参数表示 Cosmos REST API 地址, rpcURI2 参数表示 Tendermint RPC 地址
*/
func NewCmClient(rpcURI, rpcURI2 string) (*CmClient, error) {

	var (
		c   = &CmClient{}
		err error
	)
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount(config.DefaultAccountPrefix, config.DefaultAccountPrefix+sdk.PrefixPublic)
	cfg.Seal()

	// 使用 insecure.NewCredentials() 创建一个新的凭证
	creds := credentials.NewTLS(nil)

	if c.Conn, err = grpc.Dial(rpcURI2, grpc.WithTransportCredentials(creds)); err != nil {
		return c, err
	}

	if c.CliHTTP, err = http.New(rpcURI); err != nil {
		return c, err
	}

	c.cdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
	c.txConfig = tx.NewTxConfig(c.cdc, tx.DefaultSignModes)

	return c, nil
}
