package pkg

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))                       // 把 Timestamp 換算成 10 進位
	payload := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{}) // 把 Block 的 Data 都放進去
	hash := sha256.Sum256(payload)                                                // 產生 Hash
	b.Hash = hash[:]                                                              // 設置 Hash
}

func NewBlock(Data []byte, PrevBlockHash []byte) *Block {
	block := &Block{
		time.Now().Unix(), // 如此才會產生先前說的 Unix 的時間戳
		[]byte(Data),
		PrevBlockHash,
		[]byte{},
		0,
	} // 創建一個 Block 類別

	pow := NewProofOfWork(block)
	nonce, hash := pow.Proof()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func CreateGenesisBlock() *Block {
	return NewBlock([]byte("Genesis Block"), []byte{})
}

func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		log.Panic(err)
	}

	return result.Bytes()
}

func DeserializeBlock(d []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Panic(err)
	}

	return &block
}
