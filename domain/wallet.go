package domain

// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWBitcoinScript.h>
import "C"
import (
	"unsafe"

	"github.com/default23/crypto-api/lib/twutil"
)

type Wallet struct {
	wallet *C.struct_TWHDWallet
	gate   Gate
}

func (w Wallet) Free() {
	defer C.TWHDWalletDelete(w.wallet)
}

func (w Wallet) GetPrivateKey() PrivateKey {
	key := C.TWHDWalletGetKeyForCoin(w.wallet, w.gate.TWCoinType())
	keyData := C.TWPrivateKeyData(key)

	return PrivateKey(keyData)
}

func (w Wallet) GetCoinAddress() CoinAddress {
	address := C.TWHDWalletGetAddressForCoin(w.wallet, w.gate.TWCoinType())
	return CoinAddress(address)
}

func NewWallet(gate Gate, seed Seed) (Wallet, error) {
	w := Wallet{gate: gate}

	empty := twutil.TWStringCreateWithGoString("")
	defer twutil.FreeTWString(empty)

	// TODO: make a wallet creation as shared func
	w.wallet = C.TWHDWalletCreateWithMnemonic(unsafe.Pointer(seed), empty)

	return w, nil
}

func LockBitcoinScriptForAddress(a CoinAddress) unsafe.Pointer {
	script := C.TWBitcoinScriptLockScriptForAddress(unsafe.Pointer(a), GateBitcoin.TWCoinType())
	scriptData := C.TWBitcoinScriptData(script)
	defer C.TWBitcoinScriptDelete(script)

	return scriptData
}
