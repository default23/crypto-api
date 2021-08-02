package main

import (
	"github.com/default23/crypto-api/application/transaction/endpoint"
	"github.com/default23/crypto-api/domain"
	"github.com/default23/crypto-api/infrastructure/logging"
	"github.com/ybbus/jsonrpc/v2"
)

func main() {
	logger := logging.New()
	c := jsonrpc.NewClient("http://localhost:8081")

	btcTx := endpoint.SignBitcoinRequest{
		SignRequest: endpoint.SignRequest{
			Gate: string(domain.GateBitcoin),
		},
		BitcoinTransactionRequest: endpoint.BitcoinTransactionRequest{
			Tx: endpoint.BitcoinTransaction{
				Utxo: []endpoint.UnspentTransaction{
					{
						Hash:     "fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f",
						Index:    0,
						Sequence: "4294967295",
						Amount:   "625000000",
					},
				},
				ToAddress:     "1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx",
				ChangeAddress: "1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU",
				ByteFee:       1,
				Amount:        "1000000",
			},
		},
	}

	res, err := c.Call("sign_transaction", btcTx)
	if err != nil {
		logger.WithError(err).Error("bitcoin sign request failed")
	}
	if res != nil {
		if res.Error != nil {
			logger.WithError(err).Error("bitcoin sign request failed")
		} else {
			logger.Info("sign BTC successful")
			logger.Infof("%+v", res.Result)
		}
	}

	ethTx := endpoint.SignEthereumRequest{
		SignRequest: endpoint.SignRequest{
			Gate: "ethereum",
		},
		EthereumTransactionRequest: endpoint.EthereumTransactionRequest{
			Tx: endpoint.EthereumTransaction{
				ChainID:   3,
				Nonce:     1,
				GasLimit:  21000,
				GasPrice:  5000000000,
				ToAddress: "0x7788944b6dcd32f8a3042b817cfe7c5588382bd3",
				Value:     133700000000000001,
			},
		},
	}

	res, err = c.Call("sign_transaction", ethTx)
	if err != nil {
		logger.WithError(err).Error("bitcoin sign request failed")
	}
	if res != nil {
		if res.Error != nil {
			logger.WithError(err).Error("bitcoin sign request failed")
		} else {
			logger.Info("sign ETH successful")
			logger.Infof("%+v", res.Result)
		}
	}
}
