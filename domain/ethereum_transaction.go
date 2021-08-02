package domain

import (
	"math/big"
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
		ChainId:    big.NewInt(int64(tx.ChainID)).Bytes(),
		Nonce:      big.NewInt(int64(tx.Nonce)).Bytes(),
		GasPrice:   big.NewInt(int64(tx.GasPrice)).Bytes(),
		GasLimit:   big.NewInt(int64(tx.GasLimit)).Bytes(),
		ToAddress:  tx.ToAddress,
		PrivateKey: twutil.TWDataGoBytes(unsafe.Pointer(w.GetPrivateKey())),
		Transaction: &ethereum.Transaction{
			TransactionOneof: &ethereum.Transaction_Transfer_{
				Transfer: &ethereum.Transaction_Transfer{
					Amount: big.NewInt(int64(tx.Value)).Bytes(),
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
