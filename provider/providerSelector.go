package provider

import (
	"net/http"

	"github.com/BellringerQuinn/blockchainNodeGateway/model"
)

type ProviderSelector interface {
	ConstructRequest(params model.Params) (*http.Request, Provider)
	DisableProviderForNetworkAndResource(provider Provider, params model.Params)
}

type Provider int

const (
	UnavailableRequest Provider = -1
	Infura             Provider = 0
	QuickNode          Provider = 1
)

var ProviderMap = map[Provider]string{
	Infura:    "Infura",
	QuickNode: "QuickNode",
}
