package pkg

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))                       // 把 Timestamp 換算成 10 進位
	payload := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) // 把 Block 的 Data 都放進去
	hash := sha256.Sum256(payload)                                                // 產生 Hash
	b.Hash = hash[:]                                                              // 設置 Hash
}

func CreateBlock(Data []byte, PrevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(), // 如此才會產生先前說的 Unix 的時間戳
		[]byte(Data),
		PrevBlockHash,
		[]byte{},
	} // 創建一個 Block 類別

	block.SetHash() // 為 Block 設置 Hash

	return block
}

func CreateGenesisBlock() *Block {
	return CreateBlock([]byte("Genesis Block"), []byte{})
}
