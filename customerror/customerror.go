package customerror

import (
	"errors"

	"github.com/BellringerQuinn/blockchainNodeGateway/model"
)

var (
	ErrInvalidNetworkParameter = errors.New("Invalid network parameter. Please use one of: " + getAllNetworkParameters())
)

func getAllNetworkParameters() string {
	params := ""
	for key := range model.NetworkMap {
		if params == "" {
			params += key
		} else {
			params += ", " + key
		}
	}
	return params
}
