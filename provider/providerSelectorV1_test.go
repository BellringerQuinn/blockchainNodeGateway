package provider

import (
	"fmt"
	"testing"
	"time"

	"github.com/BellringerQuinn/blockchainNodeGateway/model"
	"github.com/BellringerQuinn/blockchainNodeGateway/provider/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockSleeper struct{}

func (s *MockSleeper) Sleep(duration time.Duration) {
	time.Sleep(1 * time.Second / 2)
}

func typeOf(a interface{}) string {
	return fmt.Sprintf("%T", a)
}

// Note that this test will fail if Infura's eth_chainId API is down on Eth network - in this case, err will not be nil
func TestConstructRequest(t *testing.T) {
	tests := []struct {
		name             string
		expectedProvider model.Provider
		expectedResult   interface{}
		expectedError    error
	}{
		{
			name:             "success",
			expectedProvider: model.Infura,
			expectedResult:   "success",
			expectedError:    nil,
		},
		{
			name:             "unavailable request",
			expectedProvider: model.UnavailableRequest,
			expectedResult:   nil,
			expectedError:    nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			helper := &mocks.SelectorHelper{}
			helper.On("SelectProvider", mock.Anything).Return(test.expectedProvider)
			providerSelector := &ProviderSelectorV1{
				helper: helper,
			}

			result, provider, err := providerSelector.ConstructRequest(model.Params{
				Network:  model.Eth,
				Resource: model.ChainID,
			})

			assert.Equal(t, typeOf(test.expectedResult), typeOf(result))
			assert.Equal(t, test.expectedProvider, provider)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestDisableProviderForParams(t *testing.T) {
	s := ProviderSelectorV1{
		sleeper: &MockSleeper{},
	}
	provider := model.Infura
	params := model.Params{
		Network:  model.Eth,
		Resource: model.ChainID,
	}
	supportedRequests[provider][params.Network][params.Resource] = true

	s.DisableProviderForParams(provider, params)

	assert.Equal(t, false, supportedRequests[provider][params.Network][params.Resource])

	time.Sleep(1 * time.Second)

	// We should be retrying now
	assert.Equal(t, true, supportedRequests[provider][params.Network][params.Resource])
}

func TestDisableProviderForNetwork(t *testing.T) {
	s := ProviderSelectorV1{
		sleeper: &MockSleeper{},
	}
	provider := model.Infura
	network := model.Eth

	for resource := range supportedRequests[provider][network] {
		supportedRequests[provider][network][resource] = true
	}

	s.DisableProviderForNetwork(provider, network)

	for resource := range supportedRequests[provider][network] {
		assert.Equal(t, false, supportedRequests[provider][network][resource])
	}

	time.Sleep(1 * time.Second)

	// We should be retrying now
	for resource := range supportedRequests[provider][network] {
		assert.Equal(t, true, supportedRequests[provider][network][resource])
	}
}
