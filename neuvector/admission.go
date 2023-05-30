package neuvector

import (
	"fmt"
	"strconv"
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

// Represents the body to patch a admission rule
type PatchAdmissionRuleBody = CreateAdmissionRuleBody

// Represents the complete body to patch a admission rule
type PatchAdmissionRuleBodyFull = CreateAdmissionRuleBodyFull

// Add a new admission rule
func (c *Client) CreateAdmissionRule(body CreateAdmissionRuleBody) (*CreateAdmissionRuleResponse, error) {
	var ret CreateAdmissionRuleResponseFull

	fullBody := CreateAdmissionRuleBodyFull{body}

	if err := c.Post("/admission/rule", fullBody, &ret); err != nil {
		return nil, err
	}

	return &ret.Rule, nil
}

// Add a new admission rule in the federation
func (c *Client) CreateFedAdmissionRule(body CreateAdmissionRuleBody) (*CreateAdmissionRuleResponse, error) {
	body.CfgType = "federal"

	return c.CreateAdmissionRule(body)
}

// Returns an admission rule with a specific `id`
func (c *Client) GetAdmissionRule(id int) (*GetAdmissionRuleResponse, error) {
	var ret GetAdmissionRuleResponse

	url := fmt.Sprintf("/admission/rule/%s", strconv.Itoa(id))

	if err := c.Get(url, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// Returns every admission rules
func (c *Client) GetAdmissionRules() (*GetAdmissionRulesResponse, error) {
	var ret GetAdmissionRulesResponse

	if err := c.Get("/admission/rules", &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// Delete an admission rule base function
func (c *Client) deleteAdmissionRule(id int, isFed bool) error {
	url := fmt.Sprintf("/admission/rule/%s", strconv.Itoa(id))

	if isFed {
		url += "?scope=fed"
	}

	return c.Delete(url, nil, nil)
}

// Delete an admission rule
func (c *Client) DeleteAdmissionRule(id int) error {
	return c.deleteAdmissionRule(id, false)
}

// Delete an admission rule
func (c *Client) DeleteFedAdmissionRule(id int) error {
	return c.deleteAdmissionRule(id, true)
}

// Patch an admission rule
func (c *Client) PatchAdmissionRule(body PatchAdmissionRuleBody) error {
	return c.Patch(
		"/admission/rule",
		PatchAdmissionRuleBodyFull{body},
		nil,
	)
}
