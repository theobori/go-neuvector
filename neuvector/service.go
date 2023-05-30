package neuvector

// Represents the required body to create a service
type CreateServiceBody struct {
	Name            string  `json:"name"`
	Domain          string  `json:"domain"`
	Comment         string  `json:"comment"`
	PolicyMode      *string `json:"policy_mode,omitempty"`
	BaselineProfile *string `json:"baseline_profile,omitempty"`
	NotScored       *bool   `json:"not_scored,omitempty"`
}

// Represents the full complete body to create a service
type CreateServiceBodyFull struct {
	Config CreateServiceBody `json:"config"`
}

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

// Patch a service configuration
func (c *Client) PatchServiceConfig(body PatchServiceConfigBody) error {
	return c.Patch(
		"/service/config",
		PatchServiceConfigBodyFull{body},
		nil,
	)
}

// Add a new service
func (c *Client) CreateService(body CreateServiceBody) error {
	return c.Post(
		"/service",
		CreateServiceBodyFull{body},
		nil,
	)
}
