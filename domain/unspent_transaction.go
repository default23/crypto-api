package domain

import "encoding/hex"

type TransactionHash string

func (t TransactionHash) Decode() ([]byte, error) {
	return hex.DecodeString(string(t))
}

// NewTransactionHash is a TransactionHash constructor.
func NewTransactionHash(h string) TransactionHash {
	return TransactionHash(h)
}

type UnspentTransaction struct {
	Hash     TransactionHash
	Index    uint32
	Sequence uint32
	Amount   int64
}
