package admission

import (
	"fmt"
	"strconv"

	"github.com/theobori/go-neuvector/client"
)

// Represents the body to create a admission rule
type getAdmissionRuleResponse struct {
	ID       int                      `json:"id"`
	Category string                   `json:"category"`
	Comment  string                   `json:"comment"`
	Criteria []AdmissionRuleCriterion `json:"criteria"`
	Disable  bool                     `json:"disable"`
	Critical bool                     `json:"criticals"`
	CfgType  string                   `json:"cfg_type"`
	RuleType string                   `json:"rule_type"`
	RuleMode string                   `json:"rule_mode"`
}

// Represents an admission rule
type AdmissionRule = getAdmissionRuleResponse

// Response type to get a single admission rule
type GetAdmissionRuleResponse struct {
	Rule AdmissionRule `json:"rule"`
}

// Response type to get every admission rules
type GetAdmissionRulesResponse struct {
	Ruleq []AdmissionRule `json:"rules"`
}

const (
	// Endpoint to get a specific registry
	GetAdmissionRuleEndpoint = "/admission/rule/%s"
	// Endpoint to get every admission rule
	GetAdmissionRulesEndpoint = "/admission/rules"
)

// Returns an admission rule with a specific `id`
func GetAdmissionRule(client *client.Client, id int) (*GetAdmissionRuleResponse, error) {
	var ret GetAdmissionRuleResponse

	url := fmt.Sprintf(GetAdmissionRuleEndpoint, strconv.Itoa(id))

	if err := client.Get(url, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// Returns every admission rules
func GetAdmissionRules(client *client.Client) (*GetAdmissionRulesResponse, error) {
	var ret GetAdmissionRulesResponse

	if err := client.Get(GetAdmissionRulesEndpoint, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
