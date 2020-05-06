package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/taobun/stress-script/provider"
)

func sendCommunityPoolSpendProposal(from secp256k1.PrivKeySecp256k1, path string) (sdk.TxResponse, error) {
	p, err := provider.NewBandProvider(nodeURI, from, chainID)
	if err != nil {
		panic(err)
	}
	proposal := distribution.CommunityPoolSpendProposal{}

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	err = json.Unmarshal(contents, &proposal)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	amount, err := sdk.ParseCoins("1000000000uband")
	if err != nil {
		return sdk.TxResponse{}, err
	}

	msg := gov.NewMsgSubmitProposal(proposal, amount, p.Sender())
	return p.SendTransaction([]sdk.Msg{msg}, 0, 100000, "", "", flags.BroadcastBlock)
}

func sendUpdateParams(from secp256k1.PrivKeySecp256k1, path string) (sdk.TxResponse, error) {
	p, err := provider.NewBandProvider(nodeURI, from, chainID)
	if err != nil {
		panic(err)
	}
	proposal := params.ParameterChangeProposal{}

	contents, err := ioutil.ReadFile(path)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	err = json.Unmarshal(contents, &proposal)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	amount, err := sdk.ParseCoins("1000000000uband")
	if err != nil {
		return sdk.TxResponse{}, err
	}

	msg := gov.NewMsgSubmitProposal(proposal, amount, p.Sender())
	return p.SendTransaction([]sdk.Msg{msg}, 0, 100000, "", "", flags.BroadcastBlock)
}
