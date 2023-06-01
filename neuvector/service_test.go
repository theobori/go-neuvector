package neuvector

import (
	"testing"
)

// Test if the client is able to create a service
func TestCreateService(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	config := CreateServiceBody{
		Name:      "testsvc",
		Comment:   "a simple test service",
		NotScored: new(bool),
	}

	*config.NotScored = false

	err = c.CreateService(config)

	if err != nil {
		t.Error(err)
	}
}

// Test if the client is able to get a service
func TestGetService(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	service, err := c.GetService("testsvc")

	if err != nil || service == nil {
		t.Error(err)
	}
}

// Test if the client is able to patch a service configuration
func TestPatchServiceConfig(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	config := PatchServiceConfigBody{
		Services:  []string{"nodes"},
		NotScored: new(bool),
	}

	*config.NotScored = true

	err = c.PatchServiceConfig(config)

	if err != nil {
		t.Error(err)
	}
}
