package provider

import (
	"testing"

	"github.com/BellringerQuinn/blockchainNodeGateway/model"
	"github.com/stretchr/testify/assert"
)

func TestSelectProvider(t *testing.T) {
	s := &selector{}
	params := model.Params{
		Network:  model.Eth,
		Resource: model.ChainID,
	}

	provider := s.SelectProvider(params)

	assert.NotEqual(t, model.UnavailableRequest, provider)
	assert.True(t, supportedRequests[provider][params.Network][params.Resource])

	for p := range supportedRequests {
		supportedRequests[p][params.Network][params.Resource] = false
	}
	provider = s.SelectProvider(params)

	assert.Equal(t, model.UnavailableRequest, provider)
}
