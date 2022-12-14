package main

import (
	"blockchain_go/block2/pkg"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Help() {
	fmt.Println("There are 3 operations:")
	fmt.Println("Type 1 for adding a new Block")
	fmt.Println("Type 2 for printing the Blockchain")
	fmt.Println("Type 3 for exiting")
}

func main() {
	fmt.Println("Welcome to our Blockchain project.")
	fmt.Println("Enter h for help")

	var (
		op string
	)

	NewBlockchain := pkg.CreateBlockchain() // 新增一個區塊鏈

	for true {
		fmt.Scanln(&op)
		if op == "h" {
			fmt.Println("Printing the help")
			Help() // 顯示使用方法
			continue
		} else if op == "1" {
			fmt.Println("Entering your data:")
			reader := bufio.NewReader(os.Stdin)
			data, _, _ := reader.ReadLine() // 讀一整行 input
			NewBlockchain.AddBlock(data)    // 利用 input 作為 data 來創建區塊鏈
		} else if op == "2" {
			for _, block := range NewBlockchain.Blocks {
				fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
				fmt.Printf("Data: %s\n", block.Data)
				fmt.Printf("Hash: %x\n", block.Hash)

				pow := pkg.NewProofOfWork(block)
				fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
				fmt.Println()
			} // 查詢資料
		} else if op == "3" {
			break
		}

		fmt.Println("Please Enter h, 1, 2, 3")
	}

}
