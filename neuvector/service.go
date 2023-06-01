package neuvector

import "fmt"

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
	Services        []string `json:"services"`
	PolicyMode      *string  `json:"policy_mode,omitempty"`
	BaselineProfile *string  `json:"baseline_profile,omitempty"`
	NotScored       *bool    `json:"not_scored,omitempty"`
}

// Represents the complete required body to patch a service configuration
type PatchServiceConfigBodyFull struct {
	Config PatchServiceConfigBody `json:"config"`
}

type IPPort struct {
	Ip   string `json:"ip"`
	Port uint16 `json:"port"`
}

type GetServiceResponse struct {
	Name            string          `json:"name"`
	Comment         string          `json:"comment"`
	PolicyMode      string          `json:"policy_mode"`
	ProfileMode     string          `json:"profile_mode"`
	NotScored       bool            `json:"not_scored"`
	Domain          string          `json:"domain"`
	PlatformRole    string          `json:"platform_role"`
	Members         []WorkloadBrief `json:"members"`
	PolicyRules     []PolicyRule    `json:"policy_rules"`
	ResponseRules   []ResponseRule  `json:"response_rules"`
	ServiceAddr     IPPort          `json:"service_addr"`
	IngressExposure bool            `json:"ingress_exposure"`
	EgressExposure  bool            `json:"egress_exposure"`
	BaselineProfile string          `json:"baseline_profile"`
	CapChangeMode   bool            `json:"cap_change_mode"`
	CapScorable     bool            `json:"cap_scorable"`
}

type GetServiceResponseFull struct {
	Service GetServiceResponse `json:"service"`
}

// Get a service
func (c *Client) GetService(name string) (*GetServiceResponseFull, error) {
	var ret GetServiceResponseFull

	if err := c.Get(fmt.Sprintf("/service/%s", name), &ret); err != nil {
		return nil, err
	}

	return &ret, nil
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
