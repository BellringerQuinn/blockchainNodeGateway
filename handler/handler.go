package handlers

import "net/http"

type Handler interface {
	ValidateNetwork(next http.Handler) http.Handler
	GetChainID(w http.ResponseWriter, r *http.Request)
	GetNetworkVersion(w http.ResponseWriter, r *http.Request)
}
