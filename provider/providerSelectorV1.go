package provider

import (
	"time"

	"github.com/BellringerQuinn/blockchainNodeGateway/customerror"
	"github.com/BellringerQuinn/blockchainNodeGateway/model"
	"github.com/ethereum/go-ethereum/rpc"
)

const timeToRetryProvider = 15 * time.Minute

type ProviderSelectorV1 struct {
	helper  SelectorHelper
	sleeper Sleeper
}

func NewProviderSelectorV1() ProviderSelector {
	return ProviderSelectorV1{
		helper:  selector{},
		sleeper: &RealSleeper{},
	}
}

type Sleeper interface {
	Sleep(time.Duration)
}

type RealSleeper struct{}

func (s *RealSleeper) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func (selector ProviderSelectorV1) ConstructRequest(params model.Params) (interface{}, model.Provider, error) {
	provider := selector.helper.SelectProvider(params)
	if provider == model.UnavailableRequest {
		return nil, model.UnavailableRequest, nil
	}

	myUrl := baseUrls[provider][params.Network]
	client, err := rpc.Dial(myUrl)
	if err != nil {
		return nil, provider, customerror.WrapErrors(customerror.ErrUnableToConnectToProviderOnGivenNetwork, err)
	}
	defer client.Close()

	var result interface{}
	err = client.Call(&result, queryParams[provider][params.Resource])
	if err != nil {
		return result, provider, customerror.WrapErrors(customerror.ErrUnableToFetchResource, err)
	}

	return result, provider, nil
}

func (selector ProviderSelectorV1) DisableProviderForParams(provider model.Provider, params model.Params) {
	supportedRequests[provider][params.Network][params.Resource] = false
	retryProviderForParams(provider, params, selector.sleeper)
}

func (selector ProviderSelectorV1) DisableProviderForNetwork(provider model.Provider, network model.Network) {
	for resource := range supportedRequests[provider][network] {
		supportedRequests[provider][network][resource] = false
		retryProviderForParams(provider, model.Params{
			Network:  network,
			Resource: resource,
		}, selector.sleeper)
	}
}

func retryProviderForParams(provider model.Provider, params model.Params, sleeper Sleeper) {
	go func() {
		sleeper.Sleep(timeToRetryProvider)
		supportedRequests[provider][params.Network][params.Resource] = true
	}()
}

var supportedRequests = map[model.Provider]map[model.Network]map[model.Resource]bool{
	model.Infura: map[model.Network]map[model.Resource]bool{
		model.Eth: map[model.Resource]bool{
			model.ChainID:        true,
			model.NetworkVersion: true,
		},
		model.Polygon: map[model.Resource]bool{
			model.ChainID:        true,
			model.NetworkVersion: true,
		},
	},
	model.QuickNode: map[model.Network]map[model.Resource]bool{
		model.Eth: map[model.Resource]bool{
			model.ChainID:        true,
			model.NetworkVersion: true,
		},
		model.Polygon: map[model.Resource]bool{
			model.ChainID:        true,
			model.NetworkVersion: true,
		},
	},
}

// *Note: for a production server, these base URLs should be imported via environment variable and should never be committed - treat these like secrets!
var baseUrls = map[model.Provider]map[model.Network]string{
	model.Infura: map[model.Network]string{
		model.Eth:     "https://mainnet.infura.io/v3/622e055f96d44e06aced405e65e1d60f",
		model.Polygon: "https://polygon-mainnet.infura.io/v3/622e055f96d44e06aced405e65e1d60f",
	},
	model.QuickNode: map[model.Network]string{
		model.Eth:     "https://soft-tame-cherry.quiknode.pro/163ee311b2d5396d5697fa55bc5d6c97ee23cdc8",
		model.Polygon: "https://wispy-cool-patina.matic.quiknode.pro/71fbf400997369edbbef773fbd18c728c269f94c",
	},
}

var queryParams = map[model.Provider]map[model.Resource]string{
	model.Infura: map[model.Resource]string{
		model.ChainID:        "eth_chainId",
		model.NetworkVersion: "net_version",
	},
	model.QuickNode: map[model.Resource]string{
		model.ChainID:        "eth_chainId",
		model.NetworkVersion: "net_version",
	},
}
