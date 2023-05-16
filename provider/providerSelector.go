package provider

import (
	"github.com/BellringerQuinn/blockchainNodeGateway/model"
)

type ProviderSelector interface {
	ConstructRequest(params model.Params) (interface{}, model.Provider, error)
	DisableProviderForParams(provider model.Provider, params model.Params)
	DisableProviderForNetwork(provider model.Provider, network model.Network)
}
