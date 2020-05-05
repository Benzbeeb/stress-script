package main

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/taobun/stress-script/provider"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func delegate(from secp256k1.PrivKeySecp256k1, validators []sdk.ValAddress, value int64) (sdk.TxResponse, error) {
	p, err := provider.NewBandProvider(nodeURI, from, chainID)
	if err != nil {
		panic(err)
	}
	msgs := make([]sdk.Msg, 0)
	for _, val := range validators {
		msgs = append(msgs, staking.MsgDelegate{
			DelegatorAddress: p.Sender(),
			ValidatorAddress: val,
			Amount:           sdk.NewCoin("uband", sdk.NewInt(value)),
		})
	}
	return p.SendTransaction(msgs, 0, 200000, "", "", flags.BroadcastBlock)
}

func undelegate(from secp256k1.PrivKeySecp256k1, validators []sdk.ValAddress, value int64) (sdk.TxResponse, error) {
	p, err := provider.NewBandProvider(nodeURI, from, chainID)
	if err != nil {
		panic(err)
	}
	msgs := make([]sdk.Msg, 0)
	for _, val := range validators {
		msgs = append(msgs, staking.MsgUndelegate{
			DelegatorAddress: p.Sender(),
			ValidatorAddress: val,
			Amount:           sdk.NewCoin("uband", sdk.NewInt(value)),
		})
	}
	return p.SendTransaction(msgs, 0, 200000, "", "", flags.BroadcastBlock)
}

func withdraw(from secp256k1.PrivKeySecp256k1, validators []sdk.ValAddress) (sdk.TxResponse, error) {
	p, err := provider.NewBandProvider(nodeURI, from, chainID)
	if err != nil {
		panic(err)
	}
	msgs := make([]sdk.Msg, 0)
	for _, val := range validators {
		msgs = append(msgs, distribution.MsgWithdrawDelegatorReward{
			DelegatorAddress: p.Sender(),
			ValidatorAddress: val,
		})
	}
	return p.SendTransaction(msgs, 0, 200000, "", "", flags.BroadcastBlock)
}

func createValidator(from secp256k1.PrivKeySecp256k1, amount sdk.Coin) (sdk.TxResponse, error) {
	p, err := provider.NewBandProvider(nodeURI, from, chainID)
	if err != nil {
		panic(err)
	}
	msgs := []sdk.Msg{
		staking.NewMsgCreateValidator(
			sdk.ValAddress(p.Sender()), from.PubKey(), amount,
			staking.NewDescription("St Tester", "", "", ""), staking.NewCommissionRates(
				sdk.NewDecWithPrec(1, 1),
				sdk.NewDecWithPrec(2, 1),
				sdk.NewDecWithPrec(1, 1),
			), amount.Amount,
		),
	}
	return p.SendTransaction(msgs, 0, 200000, "Create new validator", "", flags.BroadcastBlock)
}
