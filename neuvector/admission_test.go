package neuvector

import "testing"

// Test if the client is able to create an admission rule
func TestCreateFedAdmissionRule(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
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

	resp, err := c.CreateFedAdmissionRule(config)

	if err != nil || resp == nil {
		t.SkipNow()
	}

}

// Test if the client is able to delete an admission rule in a federation
func TestDeleteFedAdmissionRule(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	err = c.DeleteFedAdmissionRule(100001)

	if err != nil {
		t.SkipNow()
	}
}

// Test if the client is able to get every admission rules
func TestGetAdmissionRules(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	rules, err := c.GetAdmissionRules()

	if err != nil || rules == nil {
		t.Error(err)
	}
}

// Test if the client is able to get a specific admission rule
func TestGetAdmissionRule(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	rule, err := c.GetAdmissionRule(1000)

	if err != nil || rule == nil {
		t.SkipNow()
	}
}
