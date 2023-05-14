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
	return &http.Request{}, UnsupportedRequest
}

func (ProviderSelectorV1) DisableProviderForNetwork(provider Provider, network model.Network) {

}
