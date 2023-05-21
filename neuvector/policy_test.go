package neuvector

import (
	"testing"
)

// Test if the client is able to authenticate
func TestGetPolicies(t *testing.T) {
	client, err := NewDefaultClient()

	if err != nil {
		t.Errorf("%s", err.Error())
	}

	policies, err := GetPolicies(client)

	if err != nil {
		t.Errorf("%s", err.Error())
	}

	client.Logger.Println(policies)
}
