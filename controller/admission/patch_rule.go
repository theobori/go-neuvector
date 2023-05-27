package admission

import "github.com/theobori/go-neuvector/client"

// Represents the body to patch a admission rule
type PatchAdmissionRuleBody = CreateAdmissionRuleBody

// Represents the complete body to patch a admission rule
type PatchAdmissionRuleBodyFull = CreateAdmissionRuleBodyFull

const (
	// Endpoint to patch a specific admission rule
	PatchAdmissionRuleEndpoint = "/admission/rule"
)

// Patch an admission rule
func PatchAdmissionRule(client *client.Client, body PatchAdmissionRuleBody) error {
	fullBody := PatchAdmissionRuleBodyFull{body}

	if err := client.Patch(PatchAdmissionRuleEndpoint, fullBody, nil); err != nil {
		return nil
	}

	return nil
}
