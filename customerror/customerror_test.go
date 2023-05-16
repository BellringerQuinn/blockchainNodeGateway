package customerror

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorContains(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
		input    error
		target   error
	}{
		{
			name:     "error matches",
			expected: true,
			input:    ErrInvalidNetworkParameter,
			target:   ErrInvalidNetworkParameter,
		},
		{
			name:     "wrapped error matches",
			expected: true,
			input:    WrapErrors(ErrUnableToConnectToProviderOnGivenNetwork, ErrInvalidNetworkParameter),
			target:   ErrInvalidNetworkParameter,
		},
		{
			name:     "wrapper error matches",
			expected: true,
			input:    WrapErrors(ErrUnableToFetchResource, ErrInvalidNetworkParameter),
			target:   ErrUnableToFetchResource,
		},
		{
			name:     "internally wrapped error matches",
			expected: true,
			input:    WrapErrors(ErrUnableToFetchResource, ErrInvalidNetworkParameter, ErrUnavailableResquest),
			target:   ErrInvalidNetworkParameter,
		},
		{
			name:     "no error matches",
			expected: false,
			input:    WrapErrors(ErrInvalidNetworkParameter, ErrUnableToConnectToProviderOnGivenNetwork, ErrUnableToFetchResource),
			target:   ErrUnableToParseResult,
		},
		{
			name:     "custom errors match",
			expected: true,
			input:    errors.New("something something random target something more"),
			target:   errors.New("target"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ErrorContains(test.input, test.target)

			assert.Equal(t, test.expected, result)
		})
	}
}

func TestWrapErrors(t *testing.T) {
	tests := []struct {
		name     string
		expected error
		inputs   []error
	}{
		{
			name:     "one error",
			expected: ErrInvalidNetworkParameter,
			inputs:   []error{ErrInvalidNetworkParameter},
		},
		{
			name:     "multiple errors",
			expected: fmt.Errorf("%w: %w: %w", ErrInvalidNetworkParameter, ErrUnableToFetchResource, ErrUnavailableResquest),
			inputs:   []error{ErrInvalidNetworkParameter, ErrUnableToFetchResource, ErrUnavailableResquest},
		},
		{
			name:     "no errors",
			expected: nil,
			inputs:   nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := WrapErrors(test.inputs...)

			if test.expected == nil {
				assert.Nil(t, result)
			} else {
				assert.Equal(t, test.expected.Error(), result.Error())
			}
		})
	}
}
