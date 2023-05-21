package scan

import (
	"fmt"

	"github.com/theobori/go-neuvector/client"
)

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
	AWSKey              *AWSKey    `json:"aws_key"`
	JFrogXray           *JFrogXray `json:"jfrog_xray"`
	GCRKey              *GCRKey    `json:"gcr_key"`
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

type GetRegistriesResponse struct {
	Registries []getRegistryResponse `json:"summarys"`
}

const (
	// Endpoint to get every registries
	GetRegistriesEndpoint = "/scan/registry"
	// Endpoint to get a specific registry
	GetRegistryEndpoint = "/scan/registry/%s"
)

// Returns a list of policies
func GetRegistries(client *client.Client) (*GetRegistriesResponse, error) {
	var ret GetRegistriesResponse

	if err := client.Get(GetRegistriesEndpoint, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func GetRegistry(client *client.Client, name string) (*getRegistryResponse, error) {
	var ret getRegistryResponse

	url := fmt.Sprintf(GetRegistryEndpoint, name)

	if err := client.Get(url, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
