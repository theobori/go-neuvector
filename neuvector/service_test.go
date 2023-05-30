package neuvector

import "testing"

// Test if the client is able to create a service
func TestCreateService(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	config := CreateServiceBody{
		Name:      "test",
		Comment:   "a simple test service",
		NotScored: new(bool),
	}

	*config.NotScored = false

	err = c.CreateService(config)

	if err != nil {
		t.Errorf(err.Error())
	}
}

// Test if the client is able to patch a service configuration
func TestPatchServiceConfig(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	config := PatchServiceConfigBody{
		Services:  new([]string),
		NotScored: new(bool),
	}

	*config.NotScored = true
	*config.Services = []string{"nodes"}

	err = c.PatchServiceConfig(config)

	if err != nil {
		t.Errorf(err.Error())
	}
}
