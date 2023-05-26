package policy

import (
	"fmt"

	"github.com/theobori/go-neuvector/client"
	"github.com/theobori/go-neuvector/util"
)

const (
	// Policy minimum ID (not included)
	PolicyMinimumID = 0
	// Policy maximum ID (not include)
	PolicyMaximumID = 10000
	// Federation policy minimum ID (not included)
	FedPolicyMinimumID = 100000
	// Federation policy maximum ID (not included)
	FedPolicyMaximumID = FedPolicyMinimumID + PolicyMaximumID
)

// Returns the policy IDs in a slice
func GetPolicyIDs(policies *GetPoliciesResponse) []int {
	var ret []int

	if policies == nil {
		return ret
	}

	for _, rule := range policies.Rules {
		ret = append(ret, rule.ID)
	}

	return ret
}

// Get policy new and available ID
func GetPolicyAvailableIDs(
	client *client.Client,
	min int,
	max int,
	amount int,
) ([]int, error) {
	var ret []int

	// Get every policies
	policies, err := GetPolicies(client)

	if err != nil {
		return ret, err
	}

	// Slice of policy ids
	policyIds := GetPolicyIDs(policies)

	for id := min; amount > 0 && id < max; id++ {
		exists, err := util.ItemExists(policyIds, id)

		if err != nil || exists {
			continue
		}

		ret = append(ret, id)

		amount--
	}

	if amount > 0 {
		return ret, fmt.Errorf("there are not enough available policy IDS")
	}

	return ret, nil
}
