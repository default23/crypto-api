package domain

// #include <TrustWalletCore/TWPublicKey.h>
import "C"
import (
	"encoding/hex"
	"errors"
	"unsafe"

	"github.com/default23/crypto-api/lib/twutil"
)

var (
	ErrPublicKeyInvalid = errors.New("public key is invalid")
)

type PublicKey unsafe.Pointer

// FreePublicKey deletes the PublicKey reference
func FreePublicKey(pk PublicKey) {
	defer twutil.FreeTWData(unsafe.Pointer(pk))
}

// NewPublicKey is a PublicKey constructor.
func NewPublicKey(key string) (PublicKey, error) {
	pubKey, _ := hex.DecodeString(key)
	pubKeyData := twutil.TWDataCreateWithGoBytes(pubKey)

	if !C.TWPublicKeyIsValid(pubKeyData, C.TWPublicKeyTypeSECP256k1) {
		twutil.FreeTWData(pubKeyData)
		return nil, ErrPublicKeyInvalid
	}

	return PublicKey(pubKeyData), nil
}
