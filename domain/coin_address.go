package domain

import (
	"unsafe"

	"github.com/default23/crypto-api/lib/twutil"
)

type CoinAddress unsafe.Pointer

func FreeCoinAddress(ca CoinAddress) {
	twutil.FreeTWString(unsafe.Pointer(ca))
}
