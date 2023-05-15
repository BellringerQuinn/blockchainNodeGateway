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
