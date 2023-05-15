package resourcefetcher

import (
	"net/http"

	"github.com/BellringerQuinn/blockchainNodeGateway/model"
)

type ResourceFetcher interface {
	FetchResource(resource model.Resource, network model.Network) (string, error)
}

type WebClientInterfacer interface {
	Do(req *http.Request) (*http.Response, error)
}
