package neuvector

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/theobori/go-neuvector/util"
)

// Response type for a single policy
type getPolicyResponse struct {
	ID                    int      `json:"id"`
	Comment               string   `json:"comment"`
	From                  string   `json:"from"`
	To                    string   `json:"to"`
	Ports                 string   `json:"ports"`
	Action                string   `json:"action"`
	Applications          []string `json:"applications"`
	Learned               bool     `json:"learned"`
	Disable               bool     `json:"disable"`
	CreatedTimestamp      int64    `json:"created_timestamp"`
	LastModifiedTimestamp int64    `json:"last_modified_timestamp"`
	CfgType               string   `json:"cfg_type"`
	Priority              int      `json:"priority"`
}

// Policy rule data structure
type PolicyRule = getPolicyResponse

// Full form for a policy
type GetPolicyResponse struct {
	Rule getPolicyResponse `json:"rule"`
}

// Response type to get every policies
type GetPoliciesResponse struct {
	Rules []getPolicyResponse `json:"rules"`
}

// Returns if 2 policy rules are equal with other unique identifiers thand ID
func (p *PolicyRule) Equal(dest *PolicyRule) bool {
	return p.From == dest.From &&
		p.To == dest.To &&
		p.Ports == dest.Ports &&
		reflect.DeepEqual(p.Applications, dest.Applications) &&
		p.Learned == dest.Learned &&
		p.Disable == dest.Disable &&
		p.Comment == dest.Comment &&
		p.CfgType == dest.CfgType

}

// Data structure to insert a policy rule after another one
type PolicyRuleInsert struct {
	After int          `json:"after,omitempty"`
	Rules []PolicyRule `json:"rules"`
}

// Data structure to move a policy rule after another one
type PolicyRuleMove struct {
	After int `json:"after,omitempty"`
	ID    int `json:"id"`
}

// Represents the complete body to patch a policy rule
type PatchPolicyBody struct {
	Rules  []PolicyRule      `json:"rules"`
	Delete []int             `json:"delete"`
	Insert *PolicyRuleInsert `json:"insert,omitempty"`
	Move   *PolicyRuleMove   `json:"move,omitempty"`
}

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

// Returns a list of policies
func (c *Client) GetPolicies() (*GetPoliciesResponse, error) {
	var ret GetPoliciesResponse

	if err := c.Get("/policy/rule", &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// Returns a list of policies
func (c *Client) GetPolicy(id int) (*GetPolicyResponse, error) {
	var ret GetPolicyResponse

	url := fmt.Sprintf("/policy/rule/%s", strconv.Itoa(id))

	if err := c.Get(url, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// Returns the policy IDs in a slice
func (p *GetPoliciesResponse) GetPolicyIDs() []int {
	var ret []int

	if p == nil {
		return ret
	}

	for _, rule := range p.Rules {
		ret = append(ret, rule.ID)
	}

	return ret
}

// Get policy new and available ID
func (c *Client) GetPolicyAvailableIDs(min int, max int, amount int) ([]int, error) {
	var ret []int

	// Get every policies
	policies, err := c.GetPolicies()

	if err != nil {
		return ret, err
	}

	// Slice of policy ids
	policyIds := policies.GetPolicyIDs()

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

// Patch a policy rule
func (c *Client) PatchPolicy(body PatchPolicyBody, isFed bool) error {
	var url string

	if isFed {
		url = "/policy/rule?scope=fed"
	} else {
		url = "/policy/rule"
	}

	return c.Patch(url, body, nil)
}
