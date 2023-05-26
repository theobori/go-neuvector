package admission

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test if the client is able to create an admission rule
func TestCreateFedAdmissionRule(t *testing.T) {
	var err error

	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	config := CreateAdmissionRuleBody{
		ID:       100001,
		RuleType: "deny",
		Category: "Kubernetes",
		Criteria: []AdmissionRuleCriterion{
			{
				Name:  "runAsRoot",
				Op:    "=",
				Path:  "runAsRoot",
				Value: "true",
			},
			{
				Name:  "runAsPrivileged",
				Op:    "=",
				Path:  "runAsPrivileged",
				Value: "true",
			},
		},
		Disable: false,
	}

	resp, err := CreateFedAdmissionRule(client, config)

	if err != nil || resp == nil {
		t.Errorf(err.Error())
	}

}
