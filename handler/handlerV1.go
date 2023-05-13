package handlers

import (
	"context"
	"net/http"

	"github.com/BellringerQuinn/blockchainNodeGateway/customerror"
	"github.com/go-chi/chi/v5"
)

type HandlerV1 struct{}

func NewHandlerV1() Handler {
	return HandlerV1{}
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
		return
	}

	w.Write([]byte("GetChainID: " + network))
}

func (h HandlerV1) GetNetworkVersion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	network, ok := ctx.Value("network").(string)
	if !ok {
		return
	}

	w.Write([]byte("GetNetworkVersion: " + network))
}

func ValidateNetwork(network string) error {
	if network == "eth" || network == "polygon" {
		return nil
	} else {
		return customerror.ErrInvalidNetworkParameter
	}
}
