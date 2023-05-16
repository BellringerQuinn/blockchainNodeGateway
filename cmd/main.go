package main

import (
	"log"
	"os"

	handlers "github.com/BellringerQuinn/blockchainNodeGateway/handler"
	"github.com/BellringerQuinn/blockchainNodeGateway/provider"
	"github.com/BellringerQuinn/blockchainNodeGateway/resourcefetcher"
	"github.com/BellringerQuinn/blockchainNodeGateway/server"
)

func main() {
	providerSelector := provider.NewProviderSelectorV1()
	var logger = log.New(os.Stdout, "", 5)

	resourceFetcher := resourcefetcher.NewResourceFetcherV1(providerSelector, logger)
	server.SetupServer(handlers.NewHandlerV1(resourceFetcher))
}
