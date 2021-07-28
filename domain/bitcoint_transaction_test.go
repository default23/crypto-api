package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBitcoinTransaction_SignInput(t *testing.T) {
	seed, _ := NewSeed("observe drum fault concert analyst old short plunge loan essence symbol invite")
	wallet, _ := NewWallet(GateBitcoin, seed)

	tests := []struct {
		name        string
		transaction BitcoinTransaction
		wantErr     error
	}{
		{
			name: "success",
			transaction: BitcoinTransaction{
				Utxo: []UnspentTransaction{
					{
						Hash:     NewTransactionHash("fff7f7881a8099afa6940d42d1e7f6362bec38171ea3edf433541db4e4ad969f"),
						Index:    0,
						Sequence: 4294967295,
						Amount:   625000000,
					},
				},
				ToAddress:     "1Bp9U1ogV3A14FMvKbRJms7ctyso4Z4Tcx",
				ChangeAddress: "1FQc5LdgGHMHEN9nwkjmz6tWkxhPpxBvBU",
				ByteFee:       1,
				Amount:        1000000,
			},
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.transaction.SignInput(wallet)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
