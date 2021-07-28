package domain

// #include <TrustWalletCore/TWMnemonic.h>
import "C"
import (
	"errors"
	"unsafe"

	"github.com/default23/crypto-api/lib/twutil"
)

var (
	ErrMnemonicInvalid = errors.New("seed mnemonic is invalid")
)

type Seed unsafe.Pointer

func FreeSeed(s Seed) {
	twutil.FreeTWString(unsafe.Pointer(s))
}

// NewSeed is a Seed constructor.
func NewSeed(s string) (Seed, error) {
	str := twutil.TWStringCreateWithGoString(s)

	if !C.TWMnemonicIsValid(str) {
		return Seed(str), ErrMnemonicInvalid
	}

	return Seed(str), nil
}
