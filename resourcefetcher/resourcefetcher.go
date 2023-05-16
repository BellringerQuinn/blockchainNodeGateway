package resourcefetcher

import (
	"github.com/BellringerQuinn/blockchainNodeGateway/model"
)

type ResourceFetcher interface {
	FetchResource(model.Params) (string, error)
}
