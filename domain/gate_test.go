package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGate(t *testing.T) {
	tests := []struct {
		name     string
		gateName string
		want     Gate
		wantErr  error
	}{
		{
			name:     "success_btc",
			gateName: "bitcoin",
			want:     GateBitcoin,
			wantErr:  nil,
		},
		{
			name:     "success_eth",
			gateName: "ethereum",
			want:     GateEthereum,
			wantErr:  nil,
		},
		{
			name:     "success_case_insensitive",
			gateName: "BITcoin",
			want:     GateBitcoin,
			wantErr:  nil,
		},
		{
			name:     "error_unknown_gate",
			gateName: "some_gate",
			wantErr:  ErrUnknownGate,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGate, gotErr := NewGate(tt.gateName)

			assert.Equal(t, tt.want, gotGate)
			assert.Equal(t, tt.wantErr, gotErr)
		})
	}
}
