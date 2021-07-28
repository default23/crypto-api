package service

// #cgo CFLAGS: -I ../../../wallet-core/include
// #cgo LDFLAGS: -L ../../../wallet-core/build -L ../../../wallet-core/build/trezor-crypto -l TrustWalletCore -l protobuf -l TrezorCrypto -lc++ -lm
import "C"

import (
	"github.com/default23/crypto-api/domain"
)

type TransactionService struct {
	seed domain.Seed
}

func NewTransactionService(seed domain.Seed) *TransactionService {
	return &TransactionService{
		seed: seed,
	}
}
