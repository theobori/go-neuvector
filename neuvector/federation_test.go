package neuvector

import "testing"

// Test to promote the NeuVector instance
func TestPromote(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	err = c.Promote(
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
		t.Error(err)
	}
}

// Test to demote the NeuVector instance
func TestDemote(t *testing.T) {
	_, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	// The cluster need to be promoted for the others test
	// err = c.Demote()

	if err != nil {
		t.Error(err)
	}
}
