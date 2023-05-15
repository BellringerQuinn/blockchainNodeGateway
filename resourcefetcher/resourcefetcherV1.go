package resourcefetcher

import (
	"io"
	"log"

	"github.com/BellringerQuinn/blockchainNodeGateway/customerror"
	"github.com/BellringerQuinn/blockchainNodeGateway/model"
	"github.com/BellringerQuinn/blockchainNodeGateway/provider"
)

type ResourceFetcherV1 struct {
	provider provider.ProviderSelector
	client   WebClientInterfacer
	log      *log.Logger
}

func NewResourceFetcherV1(provider provider.ProviderSelector, client WebClientInterfacer, log *log.Logger) ResourceFetcher {
	return ResourceFetcherV1{
		provider: provider,
		client:   client,
		log:      log,
	}
}

func (fetcher ResourceFetcherV1) FetchResource(params model.Params) (string, error) {
	req, myProvider := fetcher.provider.ConstructRequest(params)
	if myProvider == provider.UnavailableRequest {
		// We have no providers that can fetch that resource on that network at the moment
		return "", customerror.ErrUnavailableResquest
	}

	response, err := fetcher.client.Do(req)
	if err != nil || response.Body == nil {
		fetcher.logProviderFailedResponse(myProvider)
		fetcher.provider.DisableProviderForNetworkAndResource(myProvider, params)
		return fetcher.FetchResource(params)
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		fetcher.logProviderFailedResponse(myProvider)
		fetcher.provider.DisableProviderForNetworkAndResource(myProvider, params)
		return fetcher.FetchResource(params)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fetcher.logProviderFailedResponse(myProvider)
		fetcher.provider.DisableProviderForNetworkAndResource(myProvider, params)
		return fetcher.FetchResource(params)
	}

	return string(body), nil
}

func (fetcher ResourceFetcherV1) logProviderFailedResponse(myProvider provider.Provider) {
	fetcher.log.Printf("Provider %s returned a failed response", provider.ProviderMap[myProvider])
}
