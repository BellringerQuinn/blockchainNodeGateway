package integrationtest

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/BellringerQuinn/blockchainNodeGateway/customerror"
	handlers "github.com/BellringerQuinn/blockchainNodeGateway/handler"
	"github.com/BellringerQuinn/blockchainNodeGateway/provider"
	"github.com/BellringerQuinn/blockchainNodeGateway/resourcefetcher"
	"github.com/BellringerQuinn/blockchainNodeGateway/server"
	"github.com/stretchr/testify/assert"
)

func TestEndpoints(t *testing.T) {
	tests := []struct {
		name                 string
		requestUrl           string
		expectedResponseBody string
		expectedResponseCode int
	}{
		{
			name:                 "eth - chainID",
			requestUrl:           "http://localhost:8080/eth/chainID",
			expectedResponseBody: "0x1",
			expectedResponseCode: http.StatusOK,
		},
		{
			name:                 "eth - networkVersion",
			requestUrl:           "http://localhost:8080/eth/networkVersion",
			expectedResponseBody: "1",
			expectedResponseCode: http.StatusOK,
		},
		{
			name:                 "polygon - chainID",
			requestUrl:           "http://localhost:8080/polygon/chainID",
			expectedResponseBody: "0x89",
			expectedResponseCode: http.StatusOK,
		},
		{
			name:                 "polygon - networkVersion",
			requestUrl:           "http://localhost:8080/polygon/networkVersion",
			expectedResponseBody: "137",
			expectedResponseCode: http.StatusOK,
		},
		{
			name:                 "not found - resource",
			requestUrl:           "http://localhost:8080/polygon/unfound",
			expectedResponseBody: "404 page not found\n",
			expectedResponseCode: http.StatusNotFound,
		},
		{
			name:                 "bad request - network",
			requestUrl:           "http://localhost:8080/unfound/chainID",
			expectedResponseBody: customerror.ErrInvalidNetworkParameter.Error(),
			expectedResponseCode: http.StatusBadRequest,
		},
	}

	providerSelector := provider.NewProviderSelectorV1()
	var logger = log.New(os.Stdout, "", 5)

	resourceFetcher := resourcefetcher.NewResourceFetcherV1(providerSelector, logger)
	go func() {
		server.SetupServer(handlers.NewHandlerV1(resourceFetcher))
	}()
	client := &http.Client{}
	time.Sleep(1 * time.Second)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, test.requestUrl, nil)
			assert.NoError(t, err)

			resp, err := client.Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, test.expectedResponseCode, resp.StatusCode)
			assert.Equal(t, test.expectedResponseBody, string(body))
		})
	}
}
