package scan

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test if the client is able to get every registries
func TestGetRegistries(t *testing.T) {
	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}

	registries, err := GetRegistries(client)

	if err != nil || registries == nil {
		t.Errorf("%s", err.Error())
	}
}

// Test if the client is able to get a specific registry
func TestGetRegistry(t *testing.T) {
	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}

	registries, err := GetRegistry(client, "docker-registry")

	if err != nil || registries == nil {
		t.Skipf("%s", err.Error())
	}
}
