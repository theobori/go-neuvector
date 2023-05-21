package client

import (
	"testing"
)

// Test if the client is able to authenticate
func TestClient(t *testing.T) {
	_, err := NewDefaultClient()

	if err != nil {
		t.Errorf("%s", err.Error())
	}
}
