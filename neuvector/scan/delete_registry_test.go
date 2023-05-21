package scan

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test if the client is able to get a specific registry
func TestDeleteRegistry(t *testing.T) {
	var err error

	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	err = DeleteRegistry(client, "docker-restry")

	if err != nil {
		t.Skipf(err.Error())
	}
}
