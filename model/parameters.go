package model

type Resource int

const (
	ChainID        Resource = 0
	NetworkVersion Resource = 1
)

type Network int

const (
	Eth     Network = 0
	Polygon Network = 1
)

var NetworkMap = map[string]Network{
	EthParam:     Eth,
	PolygonParam: Polygon,
}

const (
	EthParam     = "eth"
	PolygonParam = "polygon"
)

type Params struct {
	Network  Network
	Resource Resource
}

type Provider int

const (
	UnavailableRequest Provider = -1
	Infura             Provider = 0
	QuickNode          Provider = 1
)

var ProviderMap = map[Provider]string{
	Infura:    "Infura",
	QuickNode: "QuickNode",
}
