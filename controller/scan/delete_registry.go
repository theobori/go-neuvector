package scan

import (
	"fmt"

	"github.com/theobori/go-neuvector/client"
)

const (
	// Endpoint to delete a specific registry
	DeleteRegistryEndpoint = "/scan/registry/%s"
)

// Delete a registry
func DeleteRegistry(client *client.Client, name string) error {
	url := fmt.Sprintf(DeleteRegistryEndpoint, name)

	return client.Delete(url, nil, nil)
}