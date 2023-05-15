package resourcefetcher

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/BellringerQuinn/blockchainNodeGateway/customerror"
	"github.com/BellringerQuinn/blockchainNodeGateway/model"
	"github.com/BellringerQuinn/blockchainNodeGateway/provider"
	"github.com/BellringerQuinn/blockchainNodeGateway/resourcefetcher/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type fetchResourceTest struct {
	name                  string
	setupProviderSelector func() provider.ProviderSelector
	setupClient           func(result string) WebClientInterfacer
	expectedResult        string
	expectedError         error
}

func TestFetchResource(t *testing.T) {
	tests := []fetchResourceTest{
		{
			name: "success",
			setupProviderSelector: func() provider.ProviderSelector {
				selector := &mocks.ProviderSelector{}
				selector.On("ConstructRequest", mock.Anything).Return(&http.Request{}, provider.Infura)
				selector.On("DisableProviderForNetworkAndResource", mock.Anything, mock.Anything).Return()
				return selector
			},
			setupClient: func(result string) WebClientInterfacer {
				client := &mocks.WebClientInterfacer{}
				resp := &http.Response{StatusCode: http.StatusOK}
				resp.Body = ioutil.NopCloser(strings.NewReader(result))
				client.On("Do", mock.Anything).Return(resp, nil)
				return client
			},
			expectedResult: "success",
			expectedError:  nil,
		},
		{
			name: "fail once retry and succeed",
			setupProviderSelector: func() provider.ProviderSelector {
				selector := &mocks.ProviderSelector{}
				selector.On("ConstructRequest", mock.Anything).Return(&http.Request{}, provider.Infura).Once()
				selector.On("ConstructRequest", mock.Anything).Return(&http.Request{}, provider.QuickNode)
				selector.On("DisableProviderForNetworkAndResource", mock.Anything, mock.Anything).Return()
				return selector
			},
			setupClient: func(result string) WebClientInterfacer {
				client := &mocks.WebClientInterfacer{}
				resp := &http.Response{StatusCode: http.StatusBadRequest}
				client.On("Do", mock.Anything).Return(resp, nil).Once()
				resp.StatusCode = http.StatusOK
				resp.Body = ioutil.NopCloser(strings.NewReader(result))
				client.On("Do", mock.Anything).Return(resp, nil)
				return client
			},
			expectedResult: "success",
			expectedError:  nil,
		},
		{
			name: "fail once retry and fail out",
			setupProviderSelector: func() provider.ProviderSelector {
				selector := &mocks.ProviderSelector{}
				selector.On("ConstructRequest", mock.Anything).Return(&http.Request{}, provider.Infura).Once()
				selector.On("ConstructRequest", mock.Anything).Return(&http.Request{}, provider.UnavailableRequest).Once()
				selector.On("DisableProviderForNetworkAndResource", mock.Anything, mock.Anything).Return()
				return selector
			},
			setupClient: func(result string) WebClientInterfacer {
				client := &mocks.WebClientInterfacer{}
				resp := &http.Response{StatusCode: http.StatusBadRequest}
				client.On("Do", mock.Anything).Return(resp, nil)
				return client
			},
			expectedResult: "",
			expectedError:  customerror.ErrUnavailableResquest,
		},
		{
			name: "fail with error from provider then retry and fail out",
			setupProviderSelector: func() provider.ProviderSelector {
				selector := &mocks.ProviderSelector{}
				selector.On("ConstructRequest", mock.Anything).Return(&http.Request{}, provider.Infura).Once()
				selector.On("ConstructRequest", mock.Anything).Return(&http.Request{}, provider.QuickNode).Once()
				selector.On("ConstructRequest", mock.Anything).Return(&http.Request{}, provider.UnavailableRequest).Once()
				selector.On("DisableProviderForNetworkAndResource", mock.Anything, mock.Anything).Return()
				return selector
			},
			setupClient: func(result string) WebClientInterfacer {
				client := &mocks.WebClientInterfacer{}
				resp := &http.Response{StatusCode: http.StatusOK}
				client.On("Do", mock.Anything).Return(resp, errors.New(""))
				return client
			},
			expectedResult: "",
			expectedError:  customerror.ErrUnavailableResquest,
		},
		{
			name: "fail right away",
			setupProviderSelector: func() provider.ProviderSelector {
				selector := &mocks.ProviderSelector{}
				selector.On("ConstructRequest", mock.Anything).Return(&http.Request{}, provider.UnavailableRequest)
				selector.On("DisableProviderForNetworkAndResource", mock.Anything, mock.Anything).Return()
				return selector
			},
			setupClient: func(result string) WebClientInterfacer {
				client := &mocks.WebClientInterfacer{}
				resp := &http.Response{StatusCode: http.StatusOK}
				client.On("Do", mock.Anything).Return(resp, errors.New(""))
				return client
			},
			expectedResult: "",
			expectedError:  customerror.ErrUnavailableResquest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var logger = log.New(os.Stdout, "", 5)
			selector := test.setupProviderSelector()
			fetcher := NewResourceFetcherV1(selector, test.setupClient(test.expectedResult), logger)

			result, err := fetcher.FetchResource(model.Params{
				Network:  model.Eth,
				Resource: model.ChainID,
			})

			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
