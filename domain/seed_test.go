package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSeed(t *testing.T) {
	tests := []struct {
		name    string
		src     string
		wantErr error
	}{
		{
			name:    "success",
			src:     "observe drum fault concert analyst old short plunge loan essence symbol invite",
			wantErr: nil,
		},
		{
			name:    "error_wrong_mnemonic",
			src:     "aslkdasfnlaknal",
			wantErr: ErrMnemonicInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSeed(tt.src)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
