# blockchainNodeGateway

A blockchain node gateway integrating with Infura and Quiknode. If a request to one of the node service provider fails, the gateway will re-route the request to the other, re-enabling the failing node service provider after a short delay.

## Getting Started

1. Install Docker and start it
2. Clone this repo
3. cd into the project's root directory
4. Build the project `docker compose -f docker-compose.yml up --build
5. Hit the proxy with the sample requests below

## Sample Requests

```bash
curl http://localhost:8080/eth/chainID

curl http://localhost:8080/eth/networkVersion

curl http://localhost:8080/polygon/chainID

curl http://localhost:8080/polygon/networkVersion
```
