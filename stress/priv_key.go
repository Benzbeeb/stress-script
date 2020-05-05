package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func genPrivKeys(filename string, n int) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("Fail to create file")
	}
	defer file.Close()
	for r := 0; r < n; r++ {
		bytes := make([]byte, 32)
		rand.Read(bytes)
		file.WriteString(hex.EncodeToString(bytes) + "\n")
	}
}

func getPrivKeys(filename string, n int) []secp256k1.PrivKeySecp256k1 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Fail to open file")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	privs := make([]secp256k1.PrivKeySecp256k1, 0)
	for i := 0; i < n && scanner.Scan(); i++ {
		var priv secp256k1.PrivKeySecp256k1
		privB, _ := hex.DecodeString(scanner.Text())
		copy(priv[:], privB)
		privs = append(privs, priv)
	}
	return privs
}

func getAddress(filename string, n int) []sdk.AccAddress {
	privs := getPrivKeys(filename, n)
	addrs := make([]sdk.AccAddress, len(privs))
	for i, priv := range privs {
		addrs[i] = sdk.AccAddress(priv.PubKey().Address().Bytes())
	}
	return addrs
}
