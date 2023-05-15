package resourcefetcher

import (
	"io"

	"github.com/BellringerQuinn/blockchainNodeGateway/customerror"
	"github.com/BellringerQuinn/blockchainNodeGateway/model"
	"github.com/BellringerQuinn/blockchainNodeGateway/provider"
)

type ResourceFetcherV1 struct {
	provider provider.ProviderSelector
	client   WebClientInterfacer
}

func NewResourceFetcherV1(provider provider.ProviderSelector, client WebClientInterfacer) ResourceFetcher {
	return ResourceFetcherV1{
		provider: provider,
		client:   client,
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
		fetcher.provider.DisableProviderForNetworkAndResource(myProvider, params)
		return fetcher.FetchResource(params)
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		fetcher.provider.DisableProviderForNetworkAndResource(myProvider, params)
		return fetcher.FetchResource(params)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fetcher.provider.DisableProviderForNetworkAndResource(myProvider, params)
		return fetcher.FetchResource(params)
	}

	return string(body), nil
}
