package provider

import (
	"github.com/BellringerQuinn/blockchainNodeGateway/model"
)

type SelectorHelper interface {
	SelectProvider(params model.Params) model.Provider
}

type selector struct{}

func (selector) SelectProvider(params model.Params) model.Provider {
	for currentProvider := range supportedRequests {
		if supported, ok := supportedRequests[currentProvider][params.Network][params.Resource]; ok && supported {
			return currentProvider
		}
	}
	return model.UnavailableRequest
}
