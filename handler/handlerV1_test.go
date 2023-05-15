package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BellringerQuinn/blockchainNodeGateway/customerror"
	"github.com/BellringerQuinn/blockchainNodeGateway/handler/mocks"
	"github.com/BellringerQuinn/blockchainNodeGateway/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type validateNetworkTest struct {
	name          string
	input         string
	expectedError error
}

func TestValidateNetwork(t *testing.T) {
	tests := []validateNetworkTest{
		{
			name:          "invalid network",
			input:         "invalid",
			expectedError: customerror.ErrInvalidNetworkParameter,
		},
	}
	tests = appendAllNetworks_ValidateNetworkTest(tests)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := ValidateNetwork(test.input)

			assert.Equal(t, test.expectedError, err)
		})
	}
}

func appendAllNetworks_ValidateNetworkTest(tests []validateNetworkTest) []validateNetworkTest {
	for key := range model.NetworkMap {
		test := validateNetworkTest{
			name:          "valid network - " + key,
			input:         key,
			expectedError: nil,
		}
		tests = append(tests, test)
	}
	return tests
}

type getChainIDTest struct {
	name              string
	network           string
	expectedResponse  string
	expectedErrorCode int
	mockError         error
}

func TestGetChainID(t *testing.T) {
	tests := []getChainIDTest{
		{
			name:              "provider returns error",
			network:           model.EthParam,
			expectedResponse:  "something - likely an error from the provider",
			expectedErrorCode: http.StatusBadGateway,
			mockError:         errors.New(""),
		},
	}
	tests = appendAllNetworks_GetChainIDTest(tests)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fetcher := &mocks.ResourceFetcher{}
			fetcher.On("FetchResource", mock.Anything).Return(test.expectedResponse, test.mockError)
			handler := NewHandlerV1(fetcher)

			ctx := context.WithValue(context.TODO(), "network", test.network)
			req := http.Request{}
			req = *req.WithContext(ctx)
			w := httptest.NewRecorder()

			handler.GetChainID(w, &req)

			assert.Equal(t, test.expectedErrorCode, w.Code)
			assert.Equal(t, test.expectedResponse, w.Body.String())
		})
	}
}

func appendAllNetworks_GetChainIDTest(tests []getChainIDTest) []getChainIDTest {
	for key := range model.NetworkMap {
		test := getChainIDTest{
			name:              "valid network - " + key,
			network:           key,
			expectedResponse:  "valid response",
			expectedErrorCode: http.StatusOK,
		}
		tests = append(tests, test)
	}
	return tests
}

type getNetworkVersionTest struct {
	name              string
	network           string
	expectedResponse  string
	expectedErrorCode int
	mockError         error
}

func TestNetworkVersion(t *testing.T) {
	tests := []getNetworkVersionTest{
		{
			name:              "provider returns error",
			network:           model.EthParam,
			expectedResponse:  "something - likely an error from the provider",
			expectedErrorCode: http.StatusBadGateway,
			mockError:         errors.New(""),
		},
	}
	tests = appendAllNetworks_GetNetworkVersionTest(tests)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fetcher := &mocks.ResourceFetcher{}
			fetcher.On("FetchResource", mock.Anything).Return(test.expectedResponse, test.mockError)
			handler := NewHandlerV1(fetcher)

			ctx := context.WithValue(context.TODO(), "network", test.network)
			req := http.Request{}
			req = *req.WithContext(ctx)
			w := httptest.NewRecorder()

			handler.GetNetworkVersion(w, &req)

			assert.Equal(t, test.expectedErrorCode, w.Code)
			assert.Equal(t, test.expectedResponse, w.Body.String())
		})
	}
}

func appendAllNetworks_GetNetworkVersionTest(tests []getNetworkVersionTest) []getNetworkVersionTest {
	for key := range model.NetworkMap {
		test := getNetworkVersionTest{
			name:              "valid network - " + key,
			network:           key,
			expectedResponse:  "valid response",
			expectedErrorCode: http.StatusOK,
		}
		tests = append(tests, test)
	}
	return tests
}
