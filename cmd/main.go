package main

import (
	"net/http"
	"time"

	handlers "github.com/BellringerQuinn/blockchainNodeGateway/handler"
	"github.com/BellringerQuinn/blockchainNodeGateway/provider"
	"github.com/BellringerQuinn/blockchainNodeGateway/resourcefetcher"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	providerSelector := provider.NewProviderSelectorV1()
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	resourceFetcher := resourcefetcher.NewResourceFetcherV1(providerSelector, client)
	setupServer(handlers.NewHandlerV1(resourceFetcher))
}

func setupServer(handler handlers.Handler) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(15 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to go"))
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	r.Route("/{network}", func(r chi.Router) {
		r.Use(handler.ValidateNetwork)
		r.Get("/chainID", handler.GetChainID)
		r.Get("/networkVersion", handler.GetNetworkVersion)
	})

	http.ListenAndServe(":8080", r)
}
