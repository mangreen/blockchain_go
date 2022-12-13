package cli

import (
	"fmt"
	"log"

	"blockchain_go/block5/pkg"
)

func (cli *CLI) listAddresses() {
	wallets, err := pkg.NewWallets()
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
