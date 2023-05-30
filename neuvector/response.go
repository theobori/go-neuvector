package neuvector

type CLUSEventCondition struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type ResponseRule struct {
	ID         uint32               `json:"id"`
	Event      string               `json:"event"`
	Comment    string               `json:"comment"`
	Group      string               `json:"group"`
	Conditions []CLUSEventCondition `json:"conditions"`
	Actions    [][]string           `json:"actions"`
	Webhooks   [][]string           `json:"webhooks"`
	Disable    bool                 `json:"disable"`
	CfgType    string               `json:"cfg_type"`
}
