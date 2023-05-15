package client

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptopb "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typestx "github.com/cosmos/cosmos-sdk/types/tx"
	txsign "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsign "github.com/cosmos/cosmos-sdk/x/auth/signing"
	txpb "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authpb "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"google.golang.org/grpc"

	"me-test/config"
)

type CmClient struct {
	Conn *grpc.ClientConn

	cdc      *codec.ProtoCodec
	txConfig client.TxConfig

	TmClient    tmservice.ServiceClient
	BankClient  bankpb.QueryClient
	AuthClient  authpb.QueryClient
	StakeClient stakepb.QueryClient
}

func init() {
	// register custom Denom
	if err := sdk.RegisterDenom(sdk.MEDenom, sdk.OneDec()); err != nil {
		fmt.Println("register denom error:", err)
		return
	}
	if err := sdk.RegisterDenom(sdk.BaseMEDenom, sdk.NewDecWithPrec(1, sdk.MEExponent)); err != nil {
		fmt.Println("register denom error:", err)
		return
	}
}

func NewCmClient(grpcAddr string) (*CmClient, error) {
	var (
		c      = &CmClient{}
		encCfg = simapp.MakeTestEncodingConfig()
		err    error
	)

	if grpcAddr == "" {
		grpcAddr = config.GRPCAddr
	}

	// create grpc connection
	if c.Conn, err = grpc.Dial(grpcAddr, grpc.WithInsecure()); err != nil {
		return c, err
	}

	c.cdc = codec.NewProtoCodec(encCfg.InterfaceRegistry)
	// Configure the default signature mode
	c.txConfig = txpb.NewTxConfig(c.cdc, txpb.DefaultSignModes)

	// create bank query client
	c.BankClient = bankpb.NewQueryClient(c.Conn)
	c.AuthClient = authpb.NewQueryClient(c.Conn)
	c.StakeClient = stakepb.NewQueryClient(c.Conn)

	return c, nil
}

func (c *CmClient) BuildTx(msg sdk.Msg, priv cryptopb.PrivKey, accSeq uint64, accNum uint64) (authsign.Tx, error) {
	var (
		txBuilder = c.txConfig.NewTxBuilder()
	)

	err := txBuilder.SetMsgs(msg)
	if err != nil {
		return nil, err
	}
	fees := sdk.NewCoins(sdk.NewInt64Coin(config.DefaultDenom, 1))
	txBuilder.SetGasLimit(uint64(flags.DefaultGasLimit))
	txBuilder.SetFeeAmount(fees)

	// First round: we gather all the signer infos. We use the "set empty signature" hack to do that.
	if err = txBuilder.SetSignatures(txsign.SignatureV2{
		PubKey: priv.PubKey(),
		Data: &txsign.SingleSignatureData{
			SignMode:  c.txConfig.SignModeHandler().DefaultMode(),
			Signature: nil,
		},
		Sequence: accSeq,
	}); err != nil {
		return nil, err
	}

	// Second round: all signer infos are set, so each signer can sign.
	signerData := authsign.SignerData{
		ChainID:       config.ChainID,
		AccountNumber: accNum,
		Sequence:      accSeq,
	}
	sigV2, err := tx.SignWithPrivKey(c.txConfig.SignModeHandler().DefaultMode(), signerData, txBuilder, priv, c.txConfig, accSeq)
	if err != nil {
		return nil, err
	}
	if err = txBuilder.SetSignatures(sigV2); err != nil {
		return nil, err
	}

	return txBuilder.GetTx(), nil
}

func (c *CmClient) Encoder(tx authsign.Tx) ([]byte, error) {
	txBytes, err := c.txConfig.TxEncoder()(tx)
	if err != nil {
		return nil, err
	}
	return txBytes, nil
}

func (c *CmClient) BroadcastTx(txBytes []byte) (*typestx.BroadcastTxResponse, error) {
	txClient := typestx.NewServiceClient(c.Conn)

	grpcRes, err := txClient.BroadcastTx(
		context.Background(),
		&typestx.BroadcastTxRequest{
			Mode:    typestx.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		fmt.Println("BroadcastTx is err:", err)
		return nil, err
	}
	return grpcRes, nil
}
