package federation

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test to promote the NeuVector instance
func TestPromote(t *testing.T) {
	var err error

	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	err = Promote(
		client,
		FederationMetadata{
			MasterRestInfo: MasterRestInfo{
				Port:   11443,
				Server: "localhost",
			},
			Name: "cluster.local",
			User: "admin",
		},
	)

	if err != nil {
		t.Errorf(err.Error())
	}
}

// Test to demote the NeuVector instance
func TestDemote(t *testing.T) {
	_, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// The cluster need to be promoted for the others test
	// err = Demote(client)

	if err != nil {
		t.Errorf(err.Error())
	}
}
