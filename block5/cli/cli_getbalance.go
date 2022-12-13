package cli

import (
	"fmt"
	"log"

	"blockchain_go/block5/pkg"
)

func (cli *CLI) getBalance(address string) {
	if !pkg.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := pkg.NewBlockchain(address)
	defer bc.CloseDB()

	balance := 0
	pubKeyHash := pkg.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := bc.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
