package domain

// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
import "C"
import (
	"unsafe"

	"github.com/default23/crypto-api/lib/twutil"
)

type PrivateKey unsafe.Pointer

func FreePrivateKey(pk PrivateKey) {
	defer C.TWDataDelete(unsafe.Pointer(pk))
}

// NewPrivateKey is a PrivateKey constructor.
func NewPrivateKey(seed Seed, gate Gate) (PrivateKey, error) {
	empty := twutil.TWStringCreateWithGoString("")
	defer twutil.FreeTWString(empty)

	// TODO: make a wallet creation as shared func
	wallet := C.TWHDWalletCreateWithMnemonic(unsafe.Pointer(seed), empty)
	defer C.TWHDWalletDelete(wallet)

	key := C.TWHDWalletGetKeyForCoin(wallet, gate.TWCoinType())
	keyData := C.TWPrivateKeyData(key)

	return PrivateKey(keyData), nil
}
