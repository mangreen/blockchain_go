package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
)

// CLI responsible for processing command line arguments
type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  blockchain -a $ADDRESS - Create a blockchain and send genesis block reward to $ADDRESS")
	fmt.Println("  wallet - Generates a new key-pair and saves it into the wallet file")
	fmt.Println("  balance -a $ADDRESS - Get balance of $ADDRESS")
	fmt.Println("  list - Lists all addresses from the wallet file")
	fmt.Println("  print - Print all the blocks of the blockchain")
	fmt.Println("  send -f $FROM -t $TO -a $AMOUNT - Send $AMOUNT of coins from $FROM address to $TO address")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments and processes commands
func (cli *CLI) Run() {
	cli.validateArgs()

	getBalanceCmd := flag.NewFlagSet("balance", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("blockchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("wallet", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("list", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("a", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("a", "", "The address to send genesis block reward to")
	sendFrom := sendCmd.String("f", "", "Source wallet address")
	sendTo := sendCmd.String("t", "", "Destination wallet address")
	sendAmount := sendCmd.Int("a", 0, "Amount to send")

	switch os.Args[1] {
	case "balance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "blockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "wallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "list":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.getBalance(*getBalanceAddress)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.createBlockchain(*createBlockchainAddress)
	}

	if createWalletCmd.Parsed() {
		cli.createWallet()
	}

	if listAddressesCmd.Parsed() {
		cli.listAddresses()
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.send(*sendFrom, *sendTo, *sendAmount)
	}
}