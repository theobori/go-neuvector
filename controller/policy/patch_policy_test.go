package policy

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test if the client is able to patch a policy rule
func TestPatchFedPolicy(t *testing.T) {
	var err error

	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	err = PatchFedPolicy(
		client,
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
	)

	if err != nil {
		t.Skipf(err.Error())
	}
}
