package scan

import (
	"github.com/theobori/go-neuvector/client"
)

// Represents the complete body to create a registry
type CreateRegistryBody struct {
	Name                string     `json:"name"`
	RegistryType        string     `json:"registry_type"`
	Registry            string     `json:"registry,omitempty"`
	Filters             []string   `json:"filters"`
	Username            *string    `json:"username,omitempty"`
	Password            *string    `json:"password,omitempty"`
	AuthToken           *string    `json:"auth_token,omitempty"`
	AuthWithToken       bool       `json:"auth_with_token"`
	RescanAfterDBUpdate bool       `json:"rescan_after_db_update"`
	ScanLayers          bool       `json:"scan_layers"`
	RepoLimit           *int       `json:"repo_limit,omitempty"`
	TagLimit            *int       `json:"tag_limit,omitempty"`
	Schedule            *Schedule  `json:"schedule,omitempty"`
	AWSKey              *AWSKey    `json:"aws_key,omitempty"`
	JFrogXray           *JFrogXray `json:"jfrog_xray,omitempty"`
	GCRKey              *GCRKey    `json:"gcr_key,omitempty"`
	JFrogMode           *string    `json:"jfrog_mode,omitempty"`
	JFrogAQL            *bool      `json:"jfrog_aql,omitempty"`
	GitlabExternalURL   *string    `json:"gitlab_external_url,omitempty"`
	GitlabPrivateToken  *string    `json:"gitlab_private_token,omitempty"`
	IBMCloudTokenURL    *string    `json:"ibm_cloud_token_url,omitempty"`
	IBMCloudAccount     *string    `json:"ibm_cloud_account,omitempty"`
	CfgType             *string    `json:"cfg_type,omitempty"`
}

// Represents the full complete body to create a registry
type CreateRegistryBodyFull struct {
	Config CreateRegistryBody `json:"config"`
}

const (
	// Endpoint to get a specific registry
	CreateRegistryEndpoint = "/scan/registry"
)

// Add a new registry to scan
func CreateRegistry(client *client.Client, body CreateRegistryBody) error {
	fullBody := CreateRegistryBodyFull{body}

	return client.Post(CreateRegistryEndpoint, fullBody, nil)
}
