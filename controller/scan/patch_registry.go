package scan

import (
	"fmt"

	"github.com/theobori/go-neuvector/client"
)

// Represents the type for a registry
type PatchRegistryBody = CreateRegistryBody

// Represents the complete type for a registry
type PatchRegistryBodyFull = CreateRegistryBodyFull

const (
	// Endpoint to patch a specific registry
	PatchRegistryEndpoint = "/scan/registry/%s"
)

// Patch an existing registry
func PatchRegistry(client *client.Client, body PatchRegistryBody, name string) error {
	fullBody := PatchRegistryBodyFull{body}

	url := fmt.Sprintf(PatchRegistryEndpoint, name)

	return client.Patch(url, fullBody, nil)
}
