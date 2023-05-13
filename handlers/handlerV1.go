package handlers

import "net/http"

type HandlerV1 struct {
}

func (h HandlerV1) GetChainID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetChainID response successful"))
}

func (h HandlerV1) GetNetworkVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetNetworkVersion response successful"))
}

func NewHandlerV1() Handler {
	return HandlerV1{}
}
