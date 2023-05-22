package admission

import (
	"fmt"
	"strconv"

	"github.com/theobori/go-neuvector/client"
)

const (
	// Endpoint to delete a specific admission rule
	DeleteAdmissionRuleEndpoint = "/admission/rule/%s"
)

// Delete an admission rule base function
func deleteAdmissionRule(client *client.Client, id int, isFed bool) error {
	url := fmt.Sprintf(DeleteAdmissionRuleEndpoint, strconv.Itoa(id))

	if isFed {
		url += "?scope=fed"
	}

	return client.Delete(url, nil, nil)
}

// Delete an admission rule
func DeleteAdmissionRule(client *client.Client, id int) error {
	return deleteAdmissionRule(client, id, false)
}

// Delete an admission rule
func DeleteFedAdmissionRule(client *client.Client, id int) error {
	return deleteAdmissionRule(client, id, true)
}
