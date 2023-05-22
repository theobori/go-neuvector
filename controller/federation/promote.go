package federation

import (
	"github.com/theobori/go-neuvector/client"
)

// Federation master rest informations
type MasterRestInfo struct {
	Port   int    `json:"port"`
	Server string `json:"server"`
}

// Federation cluster metadata
type FederationMetadata struct {
	MasterRestInfo MasterRestInfo `json:"master_rest_info"`
	Name           string         `json:"name"`
	User           string         `json:"user"`
}

const (
	// Endpoint to promote
	PromoteEndpoint = "/fed/promote"
	// Endpoint to demote
	DemoteEndpoint = "/fed/demote"
)

// Promote a NeuVector instance as master
func Promote(client *client.Client, body FederationMetadata) error {
	return client.Post(PromoteEndpoint, body, nil)
}

// Demote a NeuVector instance as master
func Demote(client *client.Client) error {
	return client.Post(DemoteEndpoint, nil, nil)
}
