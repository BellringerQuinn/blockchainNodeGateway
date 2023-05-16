package resourcefetcher

import (
	"log"
	"net/http"
	"os"
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
	expectedResult        string
	expectedError         error
}

func TestFetchResource(t *testing.T) {
	tests := []fetchResourceTest{
		{
			name: "success",
			setupProviderSelector: func() provider.ProviderSelector {
				selector := &mocks.ProviderSelector{}
				selector.On("ConstructRequest", mock.Anything).Return("success", model.Infura, nil)
				return selector
			},
			expectedResult: "success",
			expectedError:  nil,
		},
		{
			name: "fail to connect to server - retry and succeed",
			setupProviderSelector: func() provider.ProviderSelector {
				selector := &mocks.ProviderSelector{}
				selector.On("ConstructRequest", mock.Anything).Return(nil, model.Infura, customerror.ErrUnableToConnectToProviderOnGivenNetwork).Once()
				selector.On("ConstructRequest", mock.Anything).Return("success", model.QuickNode, nil)
				selector.On("DisableProviderForNetwork", mock.Anything, mock.Anything).Return()
				return selector
			},
			expectedResult: "success",
			expectedError:  nil,
		},
		{
			name: "fail to fetch resource - retry and succeed",
			setupProviderSelector: func() provider.ProviderSelector {
				selector := &mocks.ProviderSelector{}
				selector.On("ConstructRequest", mock.Anything).Return(nil, model.Infura, customerror.ErrUnableToFetchResource).Once()
				selector.On("ConstructRequest", mock.Anything).Return("success", model.QuickNode, nil)
				selector.On("DisableProviderForParams", mock.Anything, mock.Anything).Return()
				return selector
			},
			expectedResult: "success",
			expectedError:  nil,
		},
		{
			name: "fail to connect to server - fail to fetch resource - retry and succeed",
			setupProviderSelector: func() provider.ProviderSelector {
				selector := &mocks.ProviderSelector{}
				selector.On("ConstructRequest", mock.Anything).Return(nil, model.Infura, customerror.ErrUnableToConnectToProviderOnGivenNetwork).Once()
				selector.On("ConstructRequest", mock.Anything).Return(nil, model.QuickNode, customerror.ErrUnableToFetchResource).Once()
				selector.On("ConstructRequest", mock.Anything).Return("success", model.QuickNode, nil)
				selector.On("DisableProviderForParams", mock.Anything, mock.Anything).Return()
				selector.On("DisableProviderForNetwork", mock.Anything, mock.Anything).Return()
				return selector
			},
			expectedResult: "success",
			expectedError:  nil,
		},
		{
			name: "fail - retry and fail",
			setupProviderSelector: func() provider.ProviderSelector {
				selector := &mocks.ProviderSelector{}
				selector.On("ConstructRequest", mock.Anything).Return(nil, model.Infura, customerror.ErrUnableToFetchResource).Once()
				selector.On("ConstructRequest", mock.Anything).Return(&http.Request{}, model.UnavailableRequest, nil).Once()
				selector.On("DisableProviderForParams", mock.Anything, mock.Anything).Return()
				return selector
			},
			expectedResult: "",
			expectedError:  customerror.ErrUnavailableResquest,
		},
		{
			name: "fail right away",
			setupProviderSelector: func() provider.ProviderSelector {
				selector := &mocks.ProviderSelector{}
				selector.On("ConstructRequest", mock.Anything).Return("something", model.UnavailableRequest, nil)
				return selector
			},
			expectedResult: "",
			expectedError:  customerror.ErrUnavailableResquest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var logger = log.New(os.Stdout, "", 5)
			selector := test.setupProviderSelector()
			fetcher := NewResourceFetcherV1(selector, logger)

			result, err := fetcher.FetchResource(model.Params{
				Network:  model.Eth,
				Resource: model.ChainID,
			})

			assert.Equal(t, test.expectedResult, result)
			assert.Equal(t, test.expectedError, err)
		})
	}
}
