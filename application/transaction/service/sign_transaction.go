package service

// #include <TrustWalletCore/TWAnySigner.h>
import "C"
import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/default23/crypto-api/application/transaction"
	"github.com/default23/crypto-api/domain"
	"github.com/default23/crypto-api/infrastructure/logging"
	"github.com/default23/crypto-api/lib/twutil"
)

func (s *TransactionService) Sign(ctx context.Context, gate domain.Gate, tx transaction.TxIO) (string, error) {
	logger := logging.MustLoggerFromContext(ctx)

	wallet, err := domain.NewWallet(gate, s.seed)
	if err != nil {
		logger.WithError(err).Errorf("sign transaction failed: create wallet failed for gate %s", gate)
		return "", errors.New("unable to create the wallet from seed")
	}

	in, err := tx.SignInput(wallet)
	if err != nil {
		logger.WithError(err).Errorf("sign transaction failed: failed to create the sign input")
		return "", errors.New("unable to create transaction signing input")
	}

	inputData := twutil.TWDataCreateWithGoBytes(in)
	defer twutil.FreeTWData(inputData)

	outputData := C.TWAnySignerSign(inputData, gate.TWCoinType())
	defer twutil.FreeTWData(outputData)

	outBytes := twutil.TWDataGoBytes(outputData)
	out, err := tx.GetSignedOutput(outBytes)
	if err != nil {
		logger.WithError(err).Error("sign transaction failed: unmarshal the signing output completes with error")
		return "", errors.New("failed to generate the signing output")
	}

	return hex.EncodeToString(out), nil
}
