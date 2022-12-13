package pkg

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"strconv"
)

const targetBits = 12
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			[]byte(strconv.FormatInt(int64(pow.block.Timestamp), 16)),
			[]byte(strconv.FormatInt(int64(targetBits), 16)),
			[]byte(strconv.FormatInt(int64(nonce), 16)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Proof() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining : \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {
		data := pow.prepareData(nonce)

		hash = sha256.Sum256(data) // 返回 data 的 SHA256 校驗和
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		//比较 hashInt 和 pow.target 的大小, hashInt < pow.target, return -1
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Printf("\n\n")

	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid := hashInt.Cmp(pow.target) == -1

	return isValid
}
