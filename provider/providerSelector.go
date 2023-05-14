package provider

import (
	"net/http"

	"github.com/BellringerQuinn/blockchainNodeGateway/model"
)

type ProviderSelector interface {
	ConstructRequest(network model.Network, resource model.Resource) (*http.Request, Provider)
	DisableProviderForNetwork(provider Provider, network model.Network)
}

type Provider int

const (
	UnsupportedRequest Provider = -1
	Infura             Provider = 0
	QuickNode          Provider = 1
)
