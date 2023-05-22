package service

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test if the client is able to patch a service configuration
func TestPatchServiceConfig(t *testing.T) {
	var err error

	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	config := PatchServiceConfigBody{
		Services:  new([]string),
		NotScored: new(bool),
	}

	*config.NotScored = false
	*config.Services = []string{"nodes"}

	err = PatchServiceConfig(client, config)

	if err != nil {
		t.Errorf(err.Error())
	}
}
