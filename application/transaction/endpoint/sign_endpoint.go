package endpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/default23/crypto-api/domain"
	kitendpoint "github.com/go-kit/kit/endpoint"

	"github.com/default23/crypto-api/application/transaction"
)

func MakeSignEndpoint(service transaction.UseCase) kitendpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		signRequest := request.(SignRequest)
		io, err := mapRequestToTxIO(signRequest)
		if err != nil {
			return SignResponse{Err: err}, nil
		}

		data, err := service.Sign(ctx, domain.Gate(signRequest.Gate), io)

		return SignResponse{Data: data, Err: err}, nil
	}
}

type SignRequest struct {
	Gate    string `json:"gate"`
	Request []byte
}

type SignBitcoinRequest struct {
	SignRequest
	BitcoinTransactionRequest
}

type SignEthereumRequest struct {
	SignRequest
	EthereumTransactionRequest
}

type BitcoinTransactionRequest struct {
	Tx BitcoinTransaction `json:"tx"`
}

type EthereumTransactionRequest struct {
	Tx EthereumTransaction `json:"tx"`
}

type UnspentTransaction struct {
	Hash     string `json:"hash"`
	Index    int    `json:"index"`
	Sequence string `json:"sequence"`
	Amount   string `json:"amount"`
}

type EthereumTransaction struct {
	ChainID   int    `json:"chainId"`
	Nonce     int    `json:"nonce"`
	GasLimit  int    `json:"gasLimit"`
	GasPrice  int    `json:"gasPrice"`
	ToAddress string `json:"toAddress"`
	Value     int    `json:"value"`
}

type BitcoinTransaction struct {
	Utxo          []UnspentTransaction `json:"utxo"`
	ToAddress     string               `json:"toAddress"`
	ChangeAddress string               `json:"ChangeAddress"`
	ByteFee       int                  `json:"byteFee"`
	Amount        string               `json:"Amount"`
}

func mapRequestToTxIO(req SignRequest) (transaction.TxIO, error) {
	switch domain.Gate(req.Gate) {
	case domain.GateBitcoin:
		btcTxRequest := BitcoinTransactionRequest{}
		err := json.Unmarshal(req.Request, &btcTxRequest)
		if err != nil {
			return nil, err
		}

		return mapBitcoinTransactionRequest(btcTxRequest)
	case domain.GateEthereum:
		ethRequest := EthereumTransactionRequest{}
		err := json.Unmarshal(req.Request, &ethRequest)
		if err != nil {
			return nil, err
		}

		return mapEthereumTransactionRequest(ethRequest)
	}

	return nil, fmt.Errorf("unknown gate: %s", req.Gate)
}

func mapBitcoinTransactionRequest(r BitcoinTransactionRequest) (transaction.TxIO, error) {
	utxo := make([]domain.UnspentTransaction, 0, len(r.Tx.Utxo))
	for _, tx := range r.Tx.Utxo {
		sequence, err := strconv.Atoi(tx.Sequence)
		if err != nil {
			return nil, err
		}

		amount, err := strconv.Atoi(tx.Amount)
		if err != nil {
			return nil, err
		}

		ut := domain.UnspentTransaction{
			Hash:     domain.NewTransactionHash(tx.Hash),
			Index:    uint32(tx.Index),
			Sequence: uint32(sequence),
			Amount:   int64(amount),
		}

		utxo = append(utxo, ut)
	}

	amount, err := strconv.Atoi(r.Tx.Amount)
	if err != nil {
		return nil, err
	}

	return &domain.BitcoinTransaction{
		Utxo:          utxo,
		ToAddress:     r.Tx.ToAddress,
		ChangeAddress: r.Tx.ChangeAddress,
		ByteFee:       int64(r.Tx.ByteFee),
		Amount:        int64(amount),
	}, nil
}

func mapEthereumTransactionRequest(r EthereumTransactionRequest) (transaction.TxIO, error) {
	return &domain.EthereumTransaction{
		ChainID:   r.Tx.ChainID,
		Nonce:     r.Tx.Nonce,
		GasLimit:  r.Tx.GasLimit,
		GasPrice:  r.Tx.GasPrice,
		ToAddress: r.Tx.ToAddress,
		Value:     r.Tx.Value,
	}, nil
}

type SignResponse struct {
	Data string `json:"data"`
	Err  error  `json:"-"`
}

func (r SignResponse) Failed() error {
	return r.Err
}
