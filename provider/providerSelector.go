package provider

import (
	"net/http"

	"github.com/BellringerQuinn/blockchainNodeGateway/model"
)

type ProviderSelector interface {
	ConstructRequest(network model.Network, resource model.Resource) (*http.Request, Provider)
	DisableProviderForNetworkAndResource(provider Provider, network model.Network, resource model.Resource)
}

type Provider int

const (
	UnavailableRequest Provider = -1
	Infura             Provider = 0
	QuickNode          Provider = 1
)
