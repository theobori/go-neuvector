package neuvector

// Federation master rest informations
type MasterRestInfo struct {
	Port   int    `json:"port"`
	Server string `json:"server"`
}

// Federation cluster metadata
type FederationMetadata struct {
	MasterRestInfo MasterRestInfo `json:"master_rest_info"`
	Name           string         `json:"name"`
	User           string         `json:"user"`
}

// Promote a NeuVector instance as master
func (c *Client) Promote(body FederationMetadata) error {
	return c.Post("/fed/promote", body, nil)
}

// Demote a NeuVector instance as master
func (c *Client) Demote() error {
	return c.Post("/fed/demote", nil, nil)
}
