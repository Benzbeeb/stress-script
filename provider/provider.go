package provider

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/tendermint/tendermint/crypto"
)

// BandProvider contains context, txBuilder, private key, and address
type BandProvider struct {
	cliCtx  context.CLIContext
	txBldr  authtypes.TxBuilder
	addr    sdk.AccAddress
	privKey crypto.PrivKey
}

// NewBandProvider creates new BandProvider create new cliCtx and txBldr
func NewBandProvider(nodeURI string, privKey crypto.PrivKey, chainID string) (BandProvider, error) {
	cdc := NewCodec()
	addr := sdk.AccAddress(privKey.PubKey().Address())
	cliCtx := NewCLIContext(nodeURI, addr).WithCodec(cdc)
	num, _, err := authtypes.NewAccountRetriever(cliCtx).GetAccountNumberSequence(addr)
	if err != nil {
		return BandProvider{}, err
	}

	return BandProvider{
		cliCtx:  cliCtx,
		txBldr:  NewTxBuilder(utils.GetTxEncoder(cdc), chainID).WithAccountNumber(num),
		addr:    addr,
		privKey: privKey,
	}, nil
}

func (provider *BandProvider) Sender() sdk.AccAddress {
	return provider.addr
}

func (provider *BandProvider) SendTransaction(
	msgs []sdk.Msg, seq, gas uint64,
	memo, fees, broadcastMode string,
) (sdk.TxResponse, error) {
	cliCtx := provider.cliCtx.WithBroadcastMode(broadcastMode)
	txBldr := provider.txBldr.WithSequence(seq).WithGas(gas).WithMemo(memo).WithFees(fees)
	return completeAndBroadcastTxCLI(cliCtx, txBldr, msgs, provider.privKey)
}
