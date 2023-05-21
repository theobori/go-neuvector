package scan

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test if the client is able to create a registry
func TestCreateRegistry(t *testing.T) {
	var err error

	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf("%s", err.Error())
		return
	}

	username := "123456"
	password := "123456"

	err = CreateRegistry(
		client,
		CreateRegistryBody{
			AuthWithToken:       false,
			Filters:             []string{"*"},
			Name:                "docker-registry",
			Password:            &password,
			Registry:            "https://registry.hub.docker.com/",
			RegistryType:        "Docker Registry",
			RescanAfterDBUpdate: true,
			ScanLayers:          false,
			Username:            &username,
		},
	)

	if err != nil {
		t.Errorf("%s", err.Error())
	}
}
