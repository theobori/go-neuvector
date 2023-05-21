package service

import (
	"github.com/theobori/go-neuvector/client"
)

// Represents the required body to patch a service configuration
type PatchServiceConfigBody struct {
	Services        *[]string `json:"services,omitempty"`
	PolicyMode      *string   `json:"policy_mode,omitempty"`
	BaselineProfile *string   `json:"baseline_profile,omitempty"`
	NotScored       *bool     `json:"not_scored,omitempty"`
}

// Represents the complete required body to patch a service configuration
type PatchServiceConfigBodyFull struct {
	Config PatchServiceConfigBody `json:"config"`
}

const (
	// Endpoint to get a specific registry
	PatchServiceConfigEndpoint = "/service/config"
)

// Patch a service configuration
func PatchServiceConfig(client *client.Client, body PatchServiceConfigBody) error {
	fullBody := PatchServiceConfigBodyFull{body}

	if err := client.Patch(PatchServiceConfigEndpoint, fullBody, nil); err != nil {
		return err
	}

	return nil
}
