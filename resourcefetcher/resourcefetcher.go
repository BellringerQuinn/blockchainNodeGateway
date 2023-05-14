package resourcefetcher

import "github.com/BellringerQuinn/blockchainNodeGateway/model"

type ResourceFetcher interface {
	FetchResource(resource model.Resource, network model.Network) (string, error)
}
