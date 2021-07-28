package transaction

import (
	"context"

	"github.com/default23/crypto-api/domain"
)

type TxInGenerator interface {
	SignInput(w domain.Wallet) ([]byte, error)
}

type TxOutGenerator interface {
	GetSignedOutput(data []byte) ([]byte, error)
}

type TxIO interface {
	TxInGenerator
	TxOutGenerator
}

// UseCase represents transaction use cases.
type UseCase interface {
	Sign(ctx context.Context, gate domain.Gate, tx TxIO) (string, error)
}
