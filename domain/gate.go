package domain

import "C"
import (
	"errors"
	"strings"
)

// #include <TrustWalletCore/TWCoinType.h>
// #include <TrustWalletCore/TWBase.h>
import "C"

// Gate is a transaction gate value-object.
type Gate string

var (
	ErrUnknownGate = errors.New("gate is unknown")
)

// Available transaction gates.
const (
	GateBitcoin  Gate = "bitcoin"
	GateEthereum Gate = "ethereum"
)

// NewGate is a Gate constructor.
func NewGate(g string) (Gate, error) {
	g = strings.ToLower(g)

	switch gate := Gate(g); gate {
	case GateBitcoin, GateEthereum:
		return gate, nil
	}

	return "", ErrUnknownGate
}

func (g Gate) TWCoinType() C.enum_TWCoinType {
	switch g {
	case GateBitcoin:
		return C.TWCoinTypeBitcoin
	case GateEthereum:
		return C.TWCoinTypeEthereum
	}

	panic("gate to coin-type is not properly implemented all gates")
}
