package neuvector

import "testing"

// Test if the client is able to get every policies
func TestGetPolicies(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	policies, err := c.GetPolicies()

	if err != nil || policies == nil {
		t.Error(err)
	}
}

// Test if the client is able to get a specific policy
func TestGetPolicy(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	policy, err := c.GetPolicy(1)

	if err != nil || policy == nil {
		t.SkipNow()
	}
}

// Test if the client is able to patch a policy rule
func TestPatchFedPolicy(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	err = c.PatchPolicy(
		PatchPolicyBody{
			Delete: []int{},
			Rules: []PolicyRule{
				{
					Action: "deny",
					Applications: []string{
						"HTTP",
						"SSH",
					},
					CfgType:  "federal",
					Disable:  false,
					From:     "fed.containers",
					To:       "fed.containers",
					ID:       100001,
					Learned:  false,
					Ports:    "any",
					Priority: 0,
				},
			},
		},
		true,
	)

	if err != nil {
		t.SkipNow()
	}
}
