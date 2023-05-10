package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	"google.golang.org/grpc"

	"me-test/config"
)

type CmClient struct {
	Conn *grpc.ClientConn

	Cdc      *codec.ProtoCodec
	TxConfig client.TxConfig

	TmClient   tmservice.ServiceClient
	BankClient bankpb.QueryClient
}

func NewCmClient(grpcAddr string) (*CmClient, error) {
	var (
		c   = &CmClient{}
		err error
	)

	if grpcAddr == "" {
		grpcAddr = config.GRPCAddr
	}

	// create grpc connection
	if c.Conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure()); err != nil {
		return c, err
	}

	c.TmClient = tmservice.NewServiceClient(c.Conn)

	c.Cdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
	// Configure the default signature mode
	c.TxConfig = tx.NewTxConfig(c.Cdc, tx.DefaultSignModes)

	// create bank query client
	c.BankClient = bankpb.NewQueryClient(c.Conn)

	return c, nil
}
