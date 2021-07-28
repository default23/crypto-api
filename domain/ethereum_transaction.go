package domain

import (
	"encoding/binary"
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
		ChainId:    toByteArray(tx.ChainID),
		Nonce:      toByteArray(tx.Nonce),
		GasPrice:   toByteArray(tx.GasPrice),
		GasLimit:   toByteArray(tx.GasLimit),
		ToAddress:  tx.ToAddress,
		PrivateKey: twutil.TWDataGoBytes(unsafe.Pointer(w.GetPrivateKey())),
		Transaction: &ethereum.Transaction{
			TransactionOneof: &ethereum.Transaction_Transfer_{
				Transfer: &ethereum.Transaction_Transfer{
					Amount: toByteArray(tx.Value),
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

func toByteArray(i int) []byte {
	var out []byte
	binary.BigEndian.PutUint32(out, uint32(i))

	return out
}
