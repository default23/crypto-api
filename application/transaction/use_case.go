package transaction

import (
	"context"

	"github.com/default23/crypto-api/domain"
)

// UseCase represents transaction use cases.
type UseCase interface {
	Sign(ctx context.Context, tx domain.Transaction)
}
