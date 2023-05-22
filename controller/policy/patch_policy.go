package policy

import (
	"github.com/theobori/go-neuvector/client"
)

// Policy rule data structure
type PolicyRule = getPolicyResponse

// Data structure to insert a policy rule after another one
type PolicyRuleInsert struct {
	After int          `json:"after,omitempty"`
	Rules []PolicyRule `json:"rules"`
}

// Data structure to move a policy rule after another one
type PolicyRuleMove struct {
	After int `json:"after,omitempty"`
	ID    int `json:"id"`
}

// Represents the complete body to patch a policy rule
type PatchPolicyBody struct {
	Rules  []PolicyRule      `json:"rules"`
	Delete []int             `json:"delete"`
	Insert *PolicyRuleInsert `json:"insert,omitempty"`
	Move   *PolicyRuleMove   `json:"move,omitempty"`
}

const (
	// Endpoint to patch a specific Policy
	PatchPolicyEndpoint = "/policy/rule"
)

// Patch a policy rule
func patchPolicy(client *client.Client, body PatchPolicyBody, isFed bool) error {
	var url string

	if isFed {
		url = PatchPolicyEndpoint + "?scope=fed"
	} else {
		url = PatchPolicyEndpoint
	}

	return client.Patch(url, body, nil)
}

// Patch a policy rule
func PatchPolicy(client *client.Client, body PatchPolicyBody) error {
	return patchPolicy(client, body, false)
}

// Patch a policy rule in the federation
func PatchFedPolicy(client *client.Client, body PatchPolicyBody) error {
	return patchPolicy(client, body, true)
}
