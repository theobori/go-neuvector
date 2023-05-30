package neuvector

import "fmt"

type Schedule struct {
	Schedule string `json:"schedule"`
	Interval int    `json:"interval"`
}

type AWSKey struct {
	ID              string `json:"id"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	Region          string `json:"region"`
}

type JFrogXray struct {
	URL      string `json:"url"`
	Enable   bool   `json:"enable"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type GCRKey struct {
	JSONKey string `json:"json_key"`
}

// Response type for a single policy
type getRegistryResponse struct {
	Name                string     `json:"name"`
	RegistryType        string     `json:"registry_type"`
	Registry            string     `json:"registry"`
	Username            string     `json:"username"`
	Password            string     `json:"password"`
	AuthToken           string     `json:"auth_token"`
	AuthWithToken       bool       `json:"auth_with_token"`
	Filters             []string   `json:"filters"`
	RescanAfterDBUpdate bool       `json:"rescan_after_db_update"`
	ScanLayers          bool       `json:"scan_layers"`
	RepoLimit           int        `json:"repo_limit"`
	TagLimit            int        `json:"tag_limit"`
	Schedule            Schedule   `json:"schedule"`
	AWSKey              *AWSKey    `json:"aws_key,omitempty"`
	JFrogXray           *JFrogXray `json:"jfrog_xray,omitempty"`
	GCRKey              *GCRKey    `json:"gcr_key,omitempty"`
	JFrogMode           string     `json:"jfrog_mode"`
	GitlabExternalURL   string     `json:"gitlab_external_url"`
	GitlabPrivateToken  string     `json:"gitlab_private_token"`
	IBMCloudTokenURL    string     `json:"ibm_cloud_token_url"`
	IBMCloudAccount     string     `json:"ibm_cloud_account"`
	Status              string     `json:"status"`
	ErrorMessage        string     `json:"error_message"`
	ErrorDetail         string     `json:"error_detail"`
	StartedAt           string     `json:"started_at,omitempty"`
	Scanned             int        `json:"scanned"`
	Scheduled           int        `json:"scheduled"`
	Scanning            int        `json:"scanning"`
	Failed              int        `json:"failed"`
	CVEDBVersion        string     `json:"cvedb_version"`
	CVEDBCreateTime     string     `json:"cvedb_create_time,omitempty"`
}

// Represents a single registry
type Registry = getRegistryResponse

// Response type to get every registries
type GetRegistriesResponse struct {
	Registries []Registry `json:"summarys"`
}

// Response type to get a single registry
type GetRegistryResponse struct {
	Registry Registry `json:"summary"`
}

// Represents the body to create a registry
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

// Represents the type for a registry
type PatchRegistryBody = CreateRegistryBody

// Represents the complete type for a registry
type PatchRegistryBodyFull = CreateRegistryBodyFull

// Patch an existing registry
func (c *Client) PatchRegistry(body PatchRegistryBody, name string) error {

	return c.Patch(
		fmt.Sprintf("/scan/registry/%s", name),
		PatchRegistryBodyFull{body},
		nil,
	)
}

// Add a new registry to scan
func (c *Client) CreateRegistry(body CreateRegistryBody) error {
	return c.Post(
		"/scan/registry",
		CreateRegistryBodyFull{body},
		nil,
	)
}

// Returns a list of registry
func (c *Client) GetRegistries() (*GetRegistriesResponse, error) {
	var ret GetRegistriesResponse

	if err := c.Get("/scan/registry", &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// Returns a registry with a specific `name`
func (c *Client) GetRegistry(name string) (*GetRegistryResponse, error) {
	var ret GetRegistryResponse

	url := fmt.Sprintf("/scan/registry/%s", name)

	if err := c.Get(url, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// Delete a registry
func (c *Client) DeleteRegistry(name string) error {
	return c.Delete(
		fmt.Sprintf("/scan/registry/%s", name),
		nil,
		nil,
	)
}
