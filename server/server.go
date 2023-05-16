package server

import (
	"net/http"
	"time"

	handlers "github.com/BellringerQuinn/blockchainNodeGateway/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupServer(handler handlers.Handler) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(15 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to the node gateway"))
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
