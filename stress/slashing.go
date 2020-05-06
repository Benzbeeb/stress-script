package main

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/taobun/stress-script/provider"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func unjail(from secp256k1.PrivKeySecp256k1) (sdk.TxResponse, error) {
	p, err := provider.NewBandProvider(nodeURI, from, chainID)
	if err != nil {
		panic(err)
	}
	return p.SendTransaction([]sdk.Msg{slashing.NewMsgUnjail(sdk.ValAddress(p.Sender()))}, 0, 200000, "", "", flags.BroadcastBlock)
}
