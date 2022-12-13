package main

import (
	"blockchain_go/block3/pkg"
)

func main() {
	bc := pkg.NewBlockchain()
	defer bc.CloseDB()

	cli := pkg.CLI{
		BC: bc,
	}
	cli.Run()
}
