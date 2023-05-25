package policy

import (
	"fmt"
	"strconv"

	"github.com/theobori/go-neuvector/client"
)

// Response type for a single policy
type getPolicyResponse struct {
	ID                    int      `json:"id"`
	Comment               string   `json:"comment"`
	From                  string   `json:"from"`
	To                    string   `json:"to"`
	Ports                 string   `json:"ports"`
	Action                string   `json:"action"`
	Applications          []string `json:"applications"`
	Learned               bool     `json:"learned"`
	Disable               bool     `json:"disable"`
	CreatedTimestamp      int64    `json:"created_timestamp"`
	LastModifiedTimestamp int64    `json:"last_modified_timestamp"`
	CfgType               string   `json:"cfg_type"`
	Priority              int      `json:"priority"`
}

// Full form for a policy
type GetPolicyResponse struct {
	Rule getPolicyResponse `json:"rule"`
}

// Response type to get every policies
type GetPoliciesResponse struct {
	Rules []getPolicyResponse `json:"rules"`
}

const (
	// Endpoint to get every policies
	GetPoliciesEndpoint = "/policy/rule"
	// Endpoint to get a specific policy
	GetPolicyEndpoint = "/policy/rule/%s"
)

// Returns a list of policies
func GetPolicies(client *client.Client) (*GetPoliciesResponse, error) {
	var ret GetPoliciesResponse

	if err := client.Get(GetPoliciesEndpoint, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// Returns a list of policies
func GetPolicy(client *client.Client, id int) (*GetPolicyResponse, error) {
	var ret GetPolicyResponse

	url := fmt.Sprintf(GetPolicyEndpoint, strconv.Itoa(id))

	if err := client.Get(url, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
