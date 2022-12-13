package cli

import (
	"fmt"
	"log"

	"blockchain_go/block5/pkg"
)

func (cli *CLI) send(from, to string, amount int) {
	if !pkg.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !pkg.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

	bc := pkg.NewBlockchain(from)
	defer bc.CloseDB()

	tx := pkg.NewUTXOTransaction(from, to, amount, bc)
	bc.MineBlock([]*pkg.Transaction{tx})
	fmt.Println("Success!")
}
