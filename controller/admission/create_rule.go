package admission

import (
	"github.com/theobori/go-neuvector/client"
)

type AdmissionRuleCriterion struct {
	Name         string  `json:"name"`
	Op           string  `json:"op"`
	Value        string  `json:"value"`
	Type         *string `json:"type,omitempty"`
	TemplateKind *string `json:"template_kind,omitempty"`
	Path         string  `json:"path,omitempty"`
	ValueType    *string `json:"value_type,omitempty"`
}

// Represents the body to create a admission rule
type CreateAdmissionRuleBody struct {
	ID       int                      `json:"id"`
	Category string                   `json:"category"`
	Comment  string                   `json:"comment"`
	Criteria []AdmissionRuleCriterion `json:"criteria,omitempty"`
	Disable  bool                     `json:"disable"`
	Actions  *[][]any                 `json:"actions,omitempty"`
	CfgType  string                   `json:"cfg_type"`
	RuleType string                   `json:"rule_type"`
	RuleMode *string                  `json:"rule_mode,omitempty"`
}

// Represents the complete body to create a admission rule
type CreateAdmissionRuleBodyFull struct {
	Config CreateAdmissionRuleBody `json:"config"`
}

// Represents the response after creating an admission rule
type CreateAdmissionRuleResponse = CreateAdmissionRuleBody

// Represents the complete response after creating an admission rule
type CreateAdmissionRuleResponseFull = struct {
	Rule CreateAdmissionRuleResponse `json:"rule"`
}

const (
	// Endpoint to create a specific admission rule
	CreateAdmissionRuleEndpoint = "/admission/rule"
)

// Add a new admission rule
func CreateAdmissionRule(client *client.Client, body CreateAdmissionRuleBody) (*CreateAdmissionRuleResponse, error) {
	var ret CreateAdmissionRuleResponseFull

	fullBody := CreateAdmissionRuleBodyFull{body}

	if err := client.Post(CreateAdmissionRuleEndpoint, fullBody, &ret); err != nil {
		return nil, err
	}

	return &ret.Rule, nil
}

// Add a new admission rule in the federation
func CreateFedAdmissionRule(client *client.Client, body CreateAdmissionRuleBody) (*CreateAdmissionRuleResponse, error) {
	body.CfgType = "federal"

	return CreateAdmissionRule(client, body)
}
