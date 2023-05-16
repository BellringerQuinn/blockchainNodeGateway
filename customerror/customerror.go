package customerror

import (
	"errors"
	"fmt"
	"strings"

	"github.com/BellringerQuinn/blockchainNodeGateway/model"
)

var (
	ErrInvalidNetworkParameter                 = errors.New("invalid network parameter. Please use one of: " + getAllNetworkParameters())
	ErrUnavailableResquest                     = errors.New("this request is currently unvailable")
	ErrUnableToGenerateRequest                 = errors.New("unable to generate request")
	ErrUnableToConnectToProviderOnGivenNetwork = errors.New("unable to connect to provider")
	ErrUnableToFetchResource                   = errors.New("unable to connect to fetch resource")
	ErrUnableToParseResult                     = errors.New("unable to connect to parse result from provider")
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

// Check if err is or contains target
func ErrorContains(err error, target error) bool {
	return strings.Contains(err.Error(), target.Error())
}

func WrapErrors(errors ...error) error {
	if len(errors) == 0 {
		return nil
	}

	err := errors[len(errors)-1]
	for i := len(errors) - 2; i >= 0; i-- {
		err = fmt.Errorf("%w: %v", errors[i], err)
	}

	return err
}
