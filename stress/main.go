package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/bandprotocol/bandchain/chain/app"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	nodeURI = "http://54.179.177.53:26657"
	chainID = "band-wenchang-testnet2"
)

func mustValAddressFromBech32(s string) sdk.ValAddress {
	v, err := sdk.ValAddressFromBech32(s)
	if err != nil {
		panic(err)
	}
	return v
}

func main() {
	config := sdk.GetConfig()
	app.SetBech32AddressPrefixesAndBip44CoinType(config)
	config.Seal()

	args := os.Args[1:]
	switch args[0] {
	case "gen-keys":
		genPrivKeys("privs.txt", 100)
		// genPrivKeys("root.txt", 1)
		// fmt.Println(getAddress("root.txt", 1)[0].String())

	case "send-beeb":
		amount, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			panic(err)
		}
		root := getPrivKeys("root.txt", 1)[0]
		beeb, _ := sdk.AccAddressFromBech32("band17sajqk0gga2s27hvv3u8vgayyp4zr5xfwhza59")
		fmt.Println(sendCoin(root, []sdk.AccAddress{beeb}, amount, 30000))
	case "many-send":
		amount, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			panic(err)
		}
		root := getPrivKeys("root.txt", 1)[0]
		privs := getAddress("privs.txt", 10)

		fmt.Println(multiSendCoin(root, privs, amount, 500000))

	case "send-back":
		amount, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			panic(err)
		}
		rootAddr := getAddress("root.txt", 1)[0]
		privs := getPrivKeys("privs.txt", 10)
		for _, priv := range privs {
			go sendCoin(priv, []sdk.AccAddress{rootAddr}, amount, 30000)
		}

	case "delegate":
		root := getPrivKeys("root.txt", 1)[0]
		addrs := getAddress("privs.txt", 1)
		fmt.Println(sendCoin(root, addrs, 1000000, 500000))

		// Delegate to switza
		validators := []sdk.ValAddress{
			mustValAddressFromBech32("bandvaloper1dzk85q35h994staarzwwnjeswrpge506splw3r"),
		}
		privs := getPrivKeys("privs.txt", 1)
		for _, priv := range privs {
			go delegate(priv, validators, 1000000)
		}
	case "delegate-fail":
		// Delegate to switza
		validators := []sdk.ValAddress{
			mustValAddressFromBech32("bandvaloper1dzk85q35h994staarzwwnjeswrpge506splw3r"),
		}
		privs := getPrivKeys("privs.txt", 10)
		for _, priv := range privs {
			go delegate(priv, validators, 100)
		}

	case "undelegate":
		validators := []sdk.ValAddress{
			mustValAddressFromBech32("bandvaloper1dzk85q35h994staarzwwnjeswrpge506splw3r"),
		}
		// Undelegate from switza
		privs := getPrivKeys("privs.txt", 10)
		for _, priv := range privs {
			go undelegate(priv, validators, 100)
		}
	case "undelegate-fail":
		validators := []sdk.ValAddress{
			mustValAddressFromBech32("bandvaloper1dzk85q35h994staarzwwnjeswrpge506splw3r"),
		}
		// Undelegate from switza
		privs := getPrivKeys("privs.txt", 10)
		for _, priv := range privs {
			go undelegate(priv, validators, 1500000)
		}

	case "withdraw":
		validators := []sdk.ValAddress{
			mustValAddressFromBech32("bandvaloper1dzk85q35h994staarzwwnjeswrpge506splw3r"),
		}
		// Withdraw reward
		privs := getPrivKeys("privs.txt", 1)
		for _, priv := range privs {
			go withdraw(priv, validators)
		}
	case "withdraw-fail":
		validators := []sdk.ValAddress{
			mustValAddressFromBech32("bandvaloper1dzk85q35h994staarzwwnjeswrpge506splw3r"),
		}
		// Withdraw reward
		privs := getPrivKeys("privs.txt", 2)
		for idx, priv := range privs {
			if idx != 0 {
				go withdraw(priv, validators)
			}
		}
	case "withdraw-val-com":
		validators := []sdk.ValAddress{
			mustValAddressFromBech32("bandvaloper1p09tl8lkl5p2g3f0jwv3evgf363el735tz8e7g"),
		}
		// Withdraw reward
		priv := getPrivKeys("privs.txt", 1)[0]

		fmt.Println(withdrawValidatorCommission(priv, validators))

	case "create-validator":
		// Create idle validator
		root := getPrivKeys("root.txt", 1)[0]
		createValidator(root, sdk.NewInt64Coin("uband", 7000000))
	case "edit-validator":
		root := getPrivKeys("root.txt", 1)[0]
		fmt.Println(editValidator(root))

	case "verify-invariant":
		root := getPrivKeys("root.txt", 1)[0]
		fmt.Println(verifyInvariant(root, 1000000))

	// case "submit-proposal":
	// 	root := getPrivKeys("root.txt", 1)[0]
	// 	addrs := getAddress("root.txt", 1)
	// 	// fmt.Println(submitProposal(root, addrs[0], 1000000))

	default:
		fmt.Println("Invalid command")

	}

	select {}
}
