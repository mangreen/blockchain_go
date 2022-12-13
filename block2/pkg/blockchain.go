package pkg

type Blockchain struct {
	Blocks []*Block
}

func (blockchain *Blockchain) AddBlock(Data []byte) {
	PrevBlock := blockchain.Blocks[len(blockchain.Blocks)-1] // 取出前一個 Block
	NewBlock := CreateBlock(Data, PrevBlock.Hash)            // 創建 Block
	blockchain.Blocks = append(blockchain.Blocks, NewBlock)  // 把新創建的 Block 接上去
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{[]*Block{CreateGenesisBlock()}}
}
