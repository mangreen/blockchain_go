package cli

import (
	"fmt"

	"blockchain_go/block5/pkg"
)

func (cli *CLI) createWallet() {
	wallets, _ := pkg.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}
