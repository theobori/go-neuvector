package scan

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test if the client is able to patch a registry
func TestPatchRegistry(t *testing.T) {
	var err error

	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	username := ""
	password := ""

	err = PatchRegistry(
		client,
		CreateRegistryBody{
			AuthWithToken:       false,
			Filters:             []string{"neuvector/*"},
			Name:                "docker-registry",
			Password:            &password,
			Registry:            "https://registry.hub.docker.com/",
			RegistryType:        "Docker Registry",
			RescanAfterDBUpdate: false,
			ScanLayers:          true,
			Username:            &username,
		},
		"docker-registry",
	)

	if err != nil {
		t.SkipNow()
	}
}
