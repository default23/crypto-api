package service

import (
	"context"

	"github.com/default23/crypto-api/domain"
)

type TransactionService struct {
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (s *TransactionService) Sign(ctx context.Context, tx domain.Transaction) {
	panic("implement me")
}
