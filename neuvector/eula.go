package neuvector

// Response type for a EULA
type getEULAReponse struct {
	Accepted bool `json:"accepted"`
}

// EULA data structure
type EULA = getEULAReponse

// Full form for a policy
type GetEULAResponse struct {
	EULA EULA `json:"eula"`
}

// Represents the body the accept EULA
type AcceptEULABody = EULA

// Represents the complete body to accept eula
type AcceptEULABodyFull = GetEULAResponse

// Accept EULA
func (c *Client) AcceptEULA(body AcceptEULABody) error {
	return c.Post(
		"/eula",
		AcceptEULABodyFull{body},
		nil,
	)
}

// Returns the EULA status
func (c *Client) GetEULA() (*GetEULAResponse, error) {
	var ret GetEULAResponse

	if err := c.Get("/eula", &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
