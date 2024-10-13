package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHAEncode(t *testing.T) {
	tests := []struct {
		name         string
		password     string
		expectedHash string
	}{
		{
			name:         "success",
			password:     "123456",
			expectedHash: ShaEncode("123456"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.expectedHash, ShaEncode(tt.password))
		})
	}
}
