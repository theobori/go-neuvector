package admission

import (
	"testing"

	"github.com/theobori/go-neuvector/client"
)

// Test if the client is able to delete an admission rule in a federation
func TestDeleteFedAdmissionRule(t *testing.T) {
	var err error

	client, err := client.NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	err = DeleteFedAdmissionRule(client, 100001)

	if err != nil {
		t.Errorf(err.Error())
	}
}
