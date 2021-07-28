package twutil

// #cgo CFLAGS: -I ../../wallet-core/include
// #cgo LDFLAGS: -L ../../wallet-core/build -L ../../wallet-core/build/trezor-crypto -l TrustWalletCore -l protobuf -l TrezorCrypto -lc++ -lm
// #include <TrustWalletCore/TWData.h>
// #include <TrustWalletCore/TWString.h>
import "C"

import (
	"encoding/hex"
	"unsafe"
)

func FreeTWString(s unsafe.Pointer) {
	C.TWStringDelete(s)
}

func FreeTWData(s unsafe.Pointer) {
	C.TWDataDelete(s)
}

// TWDataGoBytes compiles C.TWData -> Go byte[]
func TWDataGoBytes(d unsafe.Pointer) []byte {
	cBytes := C.TWDataBytes(d)
	cSize := C.TWDataSize(d)
	return C.GoBytes(unsafe.Pointer(cBytes), C.int(cSize))
}

// TWDataCreateWithGoBytes compiles Go byte[] -> C.TWData
func TWDataCreateWithGoBytes(d []byte) unsafe.Pointer {
	cBytes := C.CBytes(d)
	defer C.free(unsafe.Pointer(cBytes))
	data := C.TWDataCreateWithBytes((*C.uchar)(cBytes), C.ulong(len(d)))
	return data
}

// TWDataHexString compiles C.TWData -> Go hex string
func TWDataHexString(d unsafe.Pointer) string {
	return hex.EncodeToString(TWDataGoBytes(d))
}

// TWStringGoString compiles C.TWString -> Go string
func TWStringGoString(s unsafe.Pointer) string {
	return C.GoString(C.TWStringUTF8Bytes(s))
}

// TWStringCreateWithGoString compiles Go string -> C.TWString
func TWStringCreateWithGoString(s string) unsafe.Pointer {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))
	str := C.TWStringCreateWithUTF8Bytes(cStr)
	return str
}
