package provider

// import (
// 	"errors"
// 	"net/http"
// 	"testing"

// 	"github.com/BellringerQuinn/blockchainNodeGateway/model"
// 	"github.com/BellringerQuinn/blockchainNodeGateway/provider/mocks"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestConstructRequest(t *testing.T) {
// 	tests := []struct {
// 		name             string
// 		expectedProvider model.Provider
// 		expectedRequest  *http.Request
// 		expectedError    error
// 	}{
// 		{
// 			name:             "success",
// 			expectedProvider: model.Infura,
// 			expectedRequest:  &http.Request{Method: "GET"},
// 			expectedError:    nil,
// 		},
// 		{
// 			name:             "unavailable request",
// 			expectedProvider: model.UnavailableRequest,
// 			expectedRequest:  &http.Request{},
// 			expectedError:    nil,
// 		},
// 		{
// 			name:             "error building request",
// 			expectedProvider: model.Infura,
// 			expectedRequest:  nil,
// 			expectedError:    errors.New(""),
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			helper := &mocks.SelectorHelper{}
// 			helper.On("SelectProvider", mock.Anything).Return(test.expectedProvider)
// 			helper.On("AssembleRequestWithBody", mock.Anything, mock.Anything).Return(test.expectedRequest, test.expectedError)
// 			providerSelector := &ProviderSelectorV1{
// 				helper: helper,
// 			}

// 			request, provider, err := providerSelector.ConstructRequest(model.Params{
// 				Network:  0,
// 				Resource: 0,
// 			})

// 			assert.Equal(t, test.expectedRequest, request)
// 			assert.Equal(t, test.expectedProvider, provider)
// 			assert.Equal(t, test.expectedError, err)
// 		})
// 	}
// }
