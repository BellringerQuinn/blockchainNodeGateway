package handlers

import (
	"context"
	"net/http"

	"github.com/BellringerQuinn/blockchainNodeGateway/customerror"
	"github.com/BellringerQuinn/blockchainNodeGateway/model"
	"github.com/BellringerQuinn/blockchainNodeGateway/resourcefetcher"
	"github.com/go-chi/chi/v5"
)

type HandlerV1 struct {
	fetcher resourcefetcher.ResourceFetcher
}

func NewHandlerV1(fetcher resourcefetcher.ResourceFetcher) Handler {
	return HandlerV1{
		fetcher: fetcher,
	}
}

func (h HandlerV1) ValidateNetwork(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		network := chi.URLParam(r, "network")
		err := ValidateNetwork(network)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		ctx := context.WithValue(r.Context(), "network", network)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h HandlerV1) GetChainID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	network, ok := ctx.Value("network").(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := h.fetcher.FetchResource(model.ChainID, model.NetworkMap[network])
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	w.Write([]byte(response))
}

func (h HandlerV1) GetNetworkVersion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	network, ok := ctx.Value("network").(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := h.fetcher.FetchResource(model.NetworkVersion, model.NetworkMap[network])
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
	}
	w.Write([]byte(response))
}

func ValidateNetwork(network string) error {
	if _, ok := model.NetworkMap[network]; ok {
		return nil
	} else {
		return customerror.ErrInvalidNetworkParameter
	}
}
