package neuvector

import "fmt"

type GroupCriteria struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Op    string `json:"op"`
}

type CreateGroupBody struct {
	Name     string          `json:"name"`
	Criteria []GroupCriteria `json:"criteria"`
	CfgType  string          `json:"cfg_type"`
}

type CreateGroupBodyFull struct {
	Config CreateGroupBody `json:"config"`
}

type ScanBrief struct {
	Status           string `json:"status"`
	High             int    `json:"high"`
	Medium           int    `json:"medium"`
	Result           string `json:"result"`
	ScannedTimestamp int64  `json:"scanned_timestamp"`
	ScannedAt        string `json:"scanned_at"`
	BaseOS           string `json:"base_os"`
	ScannerVersion   string `json:"scanner_version"`
	CVEDbCreateTime  string `json:"cvedb_create_time"`
}

type WorkloadBrief struct {
	ID                 string    `json:"id"`
	Name               string    `json:"name"`
	DisplayName        string    `json:"display_name"`
	PodName            string    `json:"pod_name"`
	Image              string    `json:"image"`
	ImageID            string    `json:"image_id"`
	PlatformRole       string    `json:"platform_role"`
	Domain             string    `json:"domain"`
	State              string    `json:"state"`
	Service            string    `json:"service"`
	Author             string    `json:"author"`
	ServiceGroup       string    `json:"service_group"`
	ShareNSWith        string    `json:"share_ns_with"`
	CapSniff           bool      `json:"cap_sniff"`
	CapQuarantine      bool      `json:"cap_quarantine"`
	CapChangeMode      bool      `json:"cap_change_mode"`
	PolicyMode         string    `json:"policy_mode"`
	ProfileMode        string    `json:"profile_mode"`
	ScanSummary        ScanBrief `json:"scan_summary"`
	Children           []string  `json:"children"`
	QuarantineReason   string    `json:"quarantine_reason"`
	ServiceMesh        bool      `json:"service_mesh"`
	ServiceMeshSidecar bool      `json:"service_mesh_sidecar"`
	Privileged         bool      `json:"privileged"`
	RunAsRoot          bool      `json:"run_as_root"`
	BaselineProfile    string    `json:"baseline_profile"`
}

type GetGroupResponse struct {
	Name           string          `json:"name"`
	Learned        bool            `json:"learned"`
	Reserved       bool            `json:"reserved"`
	PolicyMode     string          `json:"policy_mode"`
	Domain         string          `json:"domain"`
	CreaterDomains []string        `json:"creater_domains"`
	Kind           string          `json:"kind"`
	PlatformRole   string          `json:"platform_role"`
	CapChangeMode  bool            `json:"cap_change_mode"`
	CfgType        string          `json:"cfg_type"`
	Criteria       []GroupCriteria `json:"criteria"`
	Members        []WorkloadBrief `json:"members"`
	PolicyRules    []PolicyRule    `json:"policy_rules"`
	ResponseRules  []ResponseRule  `json:"response_rules"`
}

type GetGroupsResponse struct {
	Name           string          `json:"name"`
	Learned        bool            `json:"learned"`
	Reserved       bool            `json:"reserved"`
	PolicyMode     string          `json:"policy_mode"`
	Domain         string          `json:"domain"`
	CreaterDomains []string        `json:"creater_domains"`
	Kind           string          `json:"kind"`
	PlatformRole   string          `json:"platform_role"`
	CapChangeMode  bool            `json:"cap_change_mode"`
	CfgType        string          `json:"cfg_type"`
	Criteria       []GroupCriteria `json:"criteria"`
	Members        []WorkloadBrief `json:"members"`
	PolicyRules    []int           `json:"policy_rules"`
	ResponseRules  []int           `json:"response_rules"`
}

type Group = GetGroupResponse

type GetGroupResponseFull struct {
	Group GetGroupResponse `json:"group"`
}

type GetGroupsResponseFull struct {
	Groups []GetGroupsResponse `json:"groups"`
}

// Represents the type for a group
type PatchGroupBody = CreateGroupBody

// Represents the complete type for a Group
type PatchGroupBodyFull = CreateGroupBodyFull

// Patch a group
func (c *Client) PatchGroup(name string, body PatchGroupBody) error {
	return c.Patch(
		fmt.Sprintf("/group/%s", name),
		PatchGroupBodyFull{body},
		nil,
	)
}

func (c *Client) GetGroups() (*GetGroupsResponseFull, error) {
	var ret GetGroupsResponseFull

	if err := c.Get("/group", &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) GetGroup(name string) (*GetGroupResponse, error) {
	var ret GetGroupResponse

	url := fmt.Sprintf("/group/%s", name)

	if err := c.Get(url, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// Add a new group
func (c *Client) CreateGroup(body CreateGroupBody) error {
	return c.Post(
		"/group",
		CreateGroupBodyFull{body},
		nil,
	)
}

// Delete a group
func (c *Client) DeleteGroup(name string) error {
	return c.Delete(
		fmt.Sprintf("/group/%s", name),
		nil,
		nil,
	)
}
