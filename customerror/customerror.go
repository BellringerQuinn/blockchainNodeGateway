package customerror

import "errors"

var (
	ErrInvalidNetworkParameter = errors.New("Invalid network parameter. Please use either 'eth' or 'polygon'.")
)
