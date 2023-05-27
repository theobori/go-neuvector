package admission

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test if the client is able to get every admission rules
func TestGetAdmissionRules(t *testing.T) {
	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	rules, err := GetAdmissionRules(client)

	if err != nil || rules == nil {
		t.Errorf(err.Error())
	}
}

// Test if the client is able to get a specific admission rule
func TestGetAdmissionRule(t *testing.T) {
	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	rule, err := GetAdmissionRule(client, 1000)

	if err != nil || rule == nil {
		t.SkipNow()
	}
}
