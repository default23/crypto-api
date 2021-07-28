package signer

// #cgo CFLAGS: -I wallet-core/include
// #cgo LDFLAGS: -L wallet-core/build -L wallet-core/build/trezor-crypto -l TrustWalletCore -l protobuf -l TrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWHDWallet.h>
// #include <TrustWalletCore/TWPrivateKey.h>
// #include <TrustWalletCore/TWPublicKey.h>
// #include <TrustWalletCore/TWBitcoinScript.h>
// #include <TrustWalletCore/TWAnySigner.h>
import "C"
