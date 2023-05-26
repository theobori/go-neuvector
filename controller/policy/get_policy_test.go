package policy

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test if the client is able to get every policies
func TestGetPolicies(t *testing.T) {
	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	policies, err := GetPolicies(client)

	if err != nil || policies == nil {
		t.Errorf(err.Error())
	}
}

// Test if the client is able to get a specific policy
func TestGetPolicy(t *testing.T) {
	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	policy, err := GetPolicy(client, 1)

	if err != nil || policy == nil {
		t.SkipNow()
	}
}
