package domain

import "C"
import (
	"unsafe"

	"google.golang.org/protobuf/proto"

	"github.com/default23/crypto-api/domain/protos/bitcoin"
	"github.com/default23/crypto-api/lib/twutil"
)

// BitcoinTransaction holds information about crypto transaction.
type BitcoinTransaction struct {
	Utxo          []UnspentTransaction
	ToAddress     string
	ChangeAddress string
	ByteFee       int64
	Amount        int64
}

func (tx *BitcoinTransaction) SignInput(w Wallet) ([]byte, error) {
	address := w.GetCoinAddress()
	defer FreeCoinAddress(address)

	scriptData := LockBitcoinScriptForAddress(address)
	defer twutil.FreeTWData(scriptData)

	utxo := make([]*bitcoin.UnspentTransaction, 0, len(tx.Utxo))

	for _, ut := range tx.Utxo {
		hash, err := ut.Hash.Decode()
		if err != nil {
			return nil, err
		}

		btcTX := &bitcoin.UnspentTransaction{
			OutPoint: &bitcoin.OutPoint{
				Hash:     hash,
				Index:    ut.Index,
				Sequence: ut.Sequence,
			},
			Amount: ut.Amount,
			Script: twutil.TWDataGoBytes(scriptData),
		}

		utxo = append(utxo, btcTX)
	}

	input := bitcoin.SigningInput{
		HashType:      1, // TWBitcoinSigHashTypeAll
		Amount:        tx.Amount,
		ByteFee:       tx.ByteFee,
		ToAddress:     tx.ToAddress,
		ChangeAddress: tx.ChangeAddress,
		PrivateKey:    [][]byte{twutil.TWDataGoBytes(unsafe.Pointer(w.GetPrivateKey()))},
		Utxo:          utxo,
		CoinType:      uint32(CoinTypeBitcoin),
	}

	inputBytes, err := proto.Marshal(&input)
	if err != nil {
		return nil, err
	}

	return inputBytes, nil
}

func (tx *BitcoinTransaction) GetSignedOutput(signed []byte) ([]byte, error) {
	var output bitcoin.SigningOutput
	err := proto.Unmarshal(signed, &output)
	if err != nil {
		return nil, err
	}

	return output.Encoded, nil
}
