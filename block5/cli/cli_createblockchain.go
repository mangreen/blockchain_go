package cli

import (
	"fmt"
	"log"

	"blockchain_go/block5/pkg"
)

func (cli *CLI) createBlockchain(address string) {
	if !pkg.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := pkg.CreateBlockchain(address)
	defer bc.CloseDB()
	fmt.Println("Done!")
}
