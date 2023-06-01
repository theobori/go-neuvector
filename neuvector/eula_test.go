package neuvector

import "testing"

// Test if the client is able to accept the EULA
func TestAcceptEULA(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	err = c.AcceptEULA(AcceptEULABody{Accepted: true})

	if err != nil {
		t.Error(err)
	}
}

// Test if the client is able to get the EULA status
func TestGetEULA(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	eula, err := c.GetEULA()

	if err != nil || eula == nil {
		t.Error(err)
	}
}
