package provider

import (
	"net/http"

	"github.com/BellringerQuinn/blockchainNodeGateway/model"
)

type ProviderSelectorV1 struct{}

func GetProviderBaseUrl(network string) string {
	return ""
}

func NewProviderSelectorV1() ProviderSelector {
	return ProviderSelectorV1{}
}

func (ProviderSelectorV1) ConstructRequest(network model.Network, resource model.Resource) (*http.Request, Provider) {
	return &http.Request{}, UnavailableRequest
}

func (ProviderSelectorV1) DisableProviderForNetworkAndResource(provider Provider, network model.Network, resource model.Resource) {

}
