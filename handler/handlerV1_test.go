package handlers

import (
	"testing"

	"github.com/BellringerQuinn/blockchainNodeGateway/customerror"
	"github.com/stretchr/testify/assert"
)

func TestValidateNetwork(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError error
	}{
		{
			name:          "invalid network",
			input:         "invalid",
			expectedError: customerror.ErrInvalidNetworkParameter,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateNetwork(test.input)

			assert.Equal(t, test.expectedError, err)
		})
	}
}
