package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	conn, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatalf("Failed to connect to infura %v", err)
	}

	contract, err := NewMultiSigWallet(common.HexToAddress("0xf1D0F9dDDc0AC93F22EE6653228f7DE8f6E2f0dc"), conn)
	if err != nil {
		log.Fatalf("Failed to interact with contarct %v", err)
	}

	log.Printf("Owners: %v", contract.GetOwners())
}
