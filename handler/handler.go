package handlers

import "net/http"

type Handler interface {
	GetChainID(w http.ResponseWriter, r *http.Request)
	GetNetworkVersion(w http.ResponseWriter, r *http.Request)
}
