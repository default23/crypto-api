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
				ChainID:   1,
				Nonce:     0,
				GasLimit:  2100,
				GasPrice:  100000,
				ToAddress: "0x17A98d2b11Dfb784e63337d2170e21cf5DD04631",
				Value:     12410,
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
