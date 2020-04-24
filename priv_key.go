package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"os"
)

func genPrivKeys(filename string, n int) {
	file, err := os.Open("filename")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for r := 0; r < n; r++ {
		bytes := make([]byte, 32)
		rand.Read(bytes)
		file.WriteString(hex.EncodeToString(bytes))
	}

}
