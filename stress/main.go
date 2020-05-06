package main

import (
	"fmt"
	"os"

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
		root := getPrivKeys("root.txt", 1)[0]
		beeb, _ := sdk.AccAddressFromBech32("band1qvw9xyk9yq47ddadtlmhq7c373673a2smxse59")
		fmt.Println(sendCoin(root, []sdk.AccAddress{beeb}, 20000000000, 30000))
	case "many-send":
		// Send 20 uband to 100 accounts
		root := getPrivKeys("root.txt", 1)[0]
		addrs := getAddress("privs.txt", 100)
		for i := 0; i < 5; i++ {
			fmt.Println(sendCoin(root, addrs[i*20:(i+1)*20], 20, 500000))
		}

	case "send-back":
		// Send 20 uband back to root
		rootAddr := getAddress("root.txt", 1)[0]
		privs := getPrivKeys("privs.txt", 100)
		for _, priv := range privs {
			go sendCoin(priv, []sdk.AccAddress{rootAddr}, 20, 30000)
		}

	case "delegate":
		// Send 1 band to 10 accounts
		root := getPrivKeys("root.txt", 1)[0]
		addrs := getAddress("privs.txt", 10)
		fmt.Println(sendCoin(root, addrs, 1000000, 500000))

		// Delegate to switza
		validators := []sdk.ValAddress{
			mustValAddressFromBech32("bandvaloper1dzk85q35h994staarzwwnjeswrpge506splw3r"),
		}
		privs := getPrivKeys("privs.txt", 10)
		for _, priv := range privs {
			go delegate(priv, validators, 1000000)
		}

	case "undelegate":
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
		privs := getPrivKeys("privs.txt", 10)
		for _, priv := range privs {
			go withdraw(priv, validators)
		}
	case "create-validator":
		// Create idle validator
		root := getPrivKeys("root.txt", 1)[0]
		createValidator(root, sdk.NewInt64Coin("uband", 20000000000))
	case "unjail":
		root := getPrivKeys("root.txt", 1)[0]
		validators := []sdk.ValAddress{
			mustValAddressFromBech32("bandvaloper1vdxjazxr9rmkxg6awtrmzrur292d3ld84kmck8"),
		}
		delegate(root, validators, 2000000)
		fmt.Println(unjail(root))
	case "community-spend":
		root := getPrivKeys("root.txt", 1)[0]
		fmt.Println(sendCommunityPoolSpendProposal(root, "proposals/community-pool-spend.json"))
	case "update-parameter":
		root := getPrivKeys("root.txt", 1)[0]
		fmt.Println(sendUpdateParams(root, "proposals/param-change.json"))
	default:
		fmt.Println("Invalid command")
	}

	select {}
}
