package client

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	cryptopb "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txpb "github.com/cosmos/cosmos-sdk/types/tx"
	txsign "github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsign "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authpb "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankpb "github.com/cosmos/cosmos-sdk/x/bank/types"
	dispb "github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakepb "github.com/cosmos/cosmos-sdk/x/staking/types"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"me-test/config"
)

type CmClient struct {
	Conn *grpc.ClientConn

	cdc      *codec.ProtoCodec
	txConfig client.TxConfig

	TxClient    txpb.ServiceClient
	TmClient    tmservice.ServiceClient
	BankClient  bankpb.QueryClient
	AuthClient  authpb.QueryClient
	StakeClient stakepb.QueryClient
	DisClient   dispb.QueryClient
}

func init() {
	// register custom Denom
	if err := sdk.RegisterDenom(sdk.MEDenom, sdk.OneDec()); err != nil {
		zap.S().Error("register denom error:", err)
		return
	}
	if err := sdk.RegisterDenom(sdk.BaseMEDenom, sdk.NewDecWithPrec(1, sdk.MEExponent)); err != nil {
		zap.S().Error("register denom error:", err)
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
	c.txConfig = authtx.NewTxConfig(c.cdc, authtx.DefaultSignModes)

	// create tx client
	c.TxClient = txpb.NewServiceClient(c.Conn)

	// create query client
	c.TmClient = tmservice.NewServiceClient(c.Conn)
	c.BankClient = bankpb.NewQueryClient(c.Conn)
	c.AuthClient = authpb.NewQueryClient(c.Conn)
	c.StakeClient = stakepb.NewQueryClient(c.Conn)
	c.DisClient = dispb.NewQueryClient(c.Conn)

	return c, nil
}

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

func (c *CmClient) BuildTx(msg sdk.Msg, priv cryptopb.PrivKey, accSeq, accNum, fee uint64) (authsign.Tx, error) {
	var (
		txBuilder = c.txConfig.NewTxBuilder()
	)

	err := txBuilder.SetMsgs(msg)
	if err != nil {
		return nil, err
	}
	fees := sdk.NewCoins(sdk.NewInt64Coin(sdk.BaseMEDenom, int64(fee)))
	txBuilder.SetGasLimit(uint64(config.DefaultGasLimit))
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

func (c *CmClient) BroadcastTx(ctx context.Context, txBytes []byte) (*txpb.BroadcastTxResponse, error) {
	grpcRes, err := c.TxClient.BroadcastTx(
		ctx,
		&txpb.BroadcastTxRequest{
			Mode:    txpb.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		zap.S().Error("BroadcastTx is err:", err)
		return nil, err
	}
	return grpcRes, nil
}

func (c *CmClient) BroadcastCheckTx(ctx context.Context, txBytes []byte) (*txpb.BroadcastTxResponse, error) {
	grpcRes, err := c.TxClient.BroadcastTx(
		ctx,
		&txpb.BroadcastTxRequest{
			Mode:    txpb.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		zap.S().Error("BroadcastTx is err:", err)
		return nil, err
	}
	return grpcRes, nil
}

func (c *CmClient) GetTx(txHash string) (*txpb.GetTxResponse, error) {
	req := &txpb.GetTxRequest{Hash: txHash}
	grpcRes, err := c.TxClient.GetTx(context.Background(), req)
	if err != nil {
		zap.S().Error("GetTx is err:", err)
		return nil, err
	}
	return grpcRes, nil
}

func (c *CmClient) SendBroadcastTx(ctx context.Context, fromPrivKey string, msg sdk.Msg, fee uint64) (*txpb.BroadcastTxResponse, error) {
	fromAccAddr, err := GetAccAddress(fromPrivKey)
	if err != nil {
		return nil, err
	}
	i, err := c.GetAccountI(ctx, fromAccAddr.String())
	if err != nil {
		return nil, err
	}
	pk256k1, _ := ConvertsAccPrivKey(fromPrivKey)
	signTx, err := c.BuildTx(msg, pk256k1, i.GetSequence(), i.GetAccountNumber(), fee)
	if err != nil {
		return nil, err
	}
	txBytes, err := c.Encoder(signTx)
	if err != nil {
		return nil, err
	}
	txRes, err := c.BroadcastTx(ctx, txBytes)
	if err != nil {
		return nil, err
	}
	zap.S().Info("SendBroadcastTx Response: ", txRes)
	return txRes, nil
}

func (c *CmClient) SendBroadcastCheckTx(ctx context.Context, fromPrivKey string, msg sdk.Msg, sequence, fee uint64) (*txpb.BroadcastTxResponse, error) {
	fromAccAddr, err := GetAccAddress(fromPrivKey)
	if err != nil {
		return nil, err
	}
	i, err := c.GetAccountI(ctx, fromAccAddr.String())
	if err != nil {
		return nil, err
	}
	pk256k1, _ := ConvertsAccPrivKey(fromPrivKey)
	signTx, err := c.BuildTx(msg, pk256k1, sequence, i.GetAccountNumber(), fee)
	if err != nil {
		return nil, err
	}
	txBytes, err := c.Encoder(signTx)
	if err != nil {
		return nil, err
	}
	txRes, err := c.BroadcastCheckTx(ctx, txBytes)
	if err != nil {
		return nil, err
	}
	zap.S().Info("SendBroadcastCheckTx Response: ", txRes)
	return txRes, nil
}
