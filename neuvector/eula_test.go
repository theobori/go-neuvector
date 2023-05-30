package neuvector

import "testing"

// Test if the client is able to accept the EULA
func TestAcceptEULA(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	err = c.AcceptEULA(AcceptEULABody{Accepted: true})

	if err != nil {
		t.Errorf(err.Error())
	}
}

// Test if the client is able to get the EULA status
func TestGetEULA(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	eula, err := c.GetEULA()

	if err != nil || eula == nil {
		t.Errorf(err.Error())
	}
}
