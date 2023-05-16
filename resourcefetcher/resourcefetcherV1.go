package resourcefetcher

import (
	"log"

	"github.com/BellringerQuinn/blockchainNodeGateway/customerror"
	"github.com/BellringerQuinn/blockchainNodeGateway/model"
	"github.com/BellringerQuinn/blockchainNodeGateway/provider"
)

type ResourceFetcherV1 struct {
	provider provider.ProviderSelector
	log      *log.Logger
}

func NewResourceFetcherV1(provider provider.ProviderSelector, log *log.Logger) ResourceFetcher {
	return ResourceFetcherV1{
		provider: provider,
		log:      log,
	}
}

func (fetcher ResourceFetcherV1) FetchResource(params model.Params) (string, error) {
	resp, myProvider, err := fetcher.provider.ConstructRequest(params)
	if myProvider == model.UnavailableRequest {
		// We have no providers that can fetch that resource on that network at the moment
		return "", customerror.ErrUnavailableResquest
	}

	if err != nil {
		fetcher.logProviderFailedResponse(myProvider, err.Error())
		if customerror.ErrorContains(err, customerror.ErrUnableToConnectToProviderOnGivenNetwork) {
			fetcher.provider.DisableProviderForNetwork(myProvider, params.Network)
		} else {
			fetcher.provider.DisableProviderForParams(myProvider, params)
		}
		return fetcher.FetchResource(params)
	}

	result, ok := resp.(string)
	if !ok {
		fetcher.logProviderFailedResponse(myProvider, customerror.ErrUnableToParseResult.Error())
		fetcher.provider.DisableProviderForParams(myProvider, params)
		return "", customerror.ErrUnableToParseResult
	}

	return result, nil
}

func (fetcher ResourceFetcherV1) logProviderFailedResponse(myProvider model.Provider, errorMessage string) {
	fetcher.log.Printf("provider %s returned a failed response: %s", model.ProviderMap[myProvider], errorMessage)
}
