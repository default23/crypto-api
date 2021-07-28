package domain

import (
	"github.com/shopspring/decimal"
)

type TransactionHash string

type UnspentTransaction struct {
	Hash     TransactionHash
	Index    int
	Sequence int
	Amount   decimal.Decimal
}
