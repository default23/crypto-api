package domain

import (
	"errors"
	"strings"
)

// Gate is a transaction gate value-object.
type Gate string

var (
	ErrUnknownGate = errors.New("gate is unknown")
)

// Available transaction gates.
const (
	GateBitcoin  = "bitcoin"
	GateEthereum = "ethereum"
)

// NewGate is a Gate constructor.
func NewGate(g string) (Gate, error) {
	g = strings.ToLower(g)

	switch g {
	case GateBitcoin, GateEthereum:
		return Gate(g), nil
	}

	return "", ErrUnknownGate
}
