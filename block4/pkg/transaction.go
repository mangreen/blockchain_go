package pkg

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// 獎勵金額
const subsidy = 10

// Transaction represents a Bitcoin transaction
type Transaction struct {
	ID    []byte
	VIns  []TxInput
	VOuts []TxOutput
}

// IsCoinbase checks whether the transaction is coinbase
func (tx Transaction) IsCoinbase() bool {
	return len(tx.VIns) == 1 && len(tx.VIns[0].TxID) == 0 && tx.VIns[0].VOutIdx == -1
}

// SetID sets ID of a transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// TxInput represents a transaction input
type TxInput struct {
	TxID      []byte
	VOutIdx   int
	ScriptSig string
}

// CanUnlockOutputWith checks whether the address initiated the transaction
func (in *TxInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

// TxOutput represents a transaction output
type TxOutput struct {
	Value        int
	ScriptPubKey string
}

// CanBeUnlockedWith checks if the output can be unlocked with the provided data
func (out *TxOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

// NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txIn := TxInput{[]byte{}, -1, data}
	txOut := TxOutput{subsidy, to}
	tx := Transaction{nil, []TxInput{txIn}, []TxOutput{txOut}}
	tx.SetID()

	return &tx
}

// NewUTxOTransaction creates a new transaction
func NewUTxOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := bc.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}

	// Build a list of inputs
	for txid, outIDs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}

		for _, outID := range outIDs {
			input := TxInput{txID, outID, from}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	outputs = append(outputs, TxOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from}) // a change
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}
