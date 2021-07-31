package domain

import (
	"encoding/hex"
	"fmt"
	"unsafe"

	"github.com/default23/crypto-api/domain/protos/ethereum"
	"github.com/default23/crypto-api/lib/twutil"
	"google.golang.org/protobuf/proto"
)

type EthereumTransaction struct {
	ChainID   int
	Nonce     int
	GasLimit  int
	GasPrice  int
	ToAddress string
	Value     int
}

func (tx EthereumTransaction) SignInput(w Wallet) ([]byte, error) {
	input := ethereum.SigningInput{
		ChainId:    toHex(tx.ChainID),
		Nonce:      toHex(tx.Nonce),
		GasPrice:   toHex(tx.GasPrice),
		GasLimit:   toHex(tx.GasLimit),
		ToAddress:  tx.ToAddress,
		PrivateKey: twutil.TWDataGoBytes(unsafe.Pointer(w.GetPrivateKey())),
		Transaction: &ethereum.Transaction{
			TransactionOneof: &ethereum.Transaction_Transfer_{
				Transfer: &ethereum.Transaction_Transfer{
					Amount: toHex(tx.Value),
				},
			},
		},
	}

	inputBytes, err := proto.Marshal(&input)
	if err != nil {
		return nil, err
	}

	return inputBytes, nil
}

func (tx *EthereumTransaction) GetSignedOutput(signed []byte) ([]byte, error) {
	var output ethereum.SigningOutput
	err := proto.Unmarshal(signed, &output)
	if err != nil {
		return nil, err
	}

	return output.Encoded, nil
}

func toHex(i int) []byte {
	// skip errors, because it will be always valid hex string
	h, _ := hex.DecodeString(fmt.Sprintf("%x", i))
	return h
}
