package domain

// #cgo CFLAGS: -I ../wallet-core/include
// #cgo LDFLAGS: -L ../wallet-core/build -L ../wallet-core/build/trezor-crypto -l TrustWalletCore -l protobuf -l TrezorCrypto -lc++ -lm
import "C"
