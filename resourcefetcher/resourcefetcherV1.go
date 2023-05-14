package resourcefetcher

import (
	"github.com/BellringerQuinn/blockchainNodeGateway/model"
	"github.com/BellringerQuinn/blockchainNodeGateway/provider"
)

type ResourceFetcherV1 struct {
	provider provider.ProviderSelector
}

func NewResourceFetcherV1(provider provider.ProviderSelector) ResourceFetcher {
	return ResourceFetcherV1{
		provider: provider,
	}
}

func (ResourceFetcherV1) FetchResource(resource model.Resource, network model.Network) (string, error) {
	return "", nil
}
