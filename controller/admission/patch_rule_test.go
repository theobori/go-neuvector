package admission

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test if the client is able to patch an admission rule
func TestPatchAdmissionRule(t *testing.T) {
	var err error

	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	config := PatchAdmissionRuleBody{
		ID:       1000,
		RuleType: "deny",
		Category: "Kubernetes",
		Criteria: []AdmissionRuleCriterion{
			{
				Name:  "runAsRoot",
				Op:    "=",
				Path:  "runAsRoot",
				Value: "true",
			},
		},
		Disable: false,
	}

	if err := PatchAdmissionRule(client, config); err != nil {
		t.SkipNow()
	}
}
