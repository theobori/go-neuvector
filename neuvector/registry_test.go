package neuvector

import "testing"

// Test if the client is able to get every registries
func TestGetRegistries(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	registries, err := c.GetRegistries()

	if err != nil || registries == nil {
		t.Errorf(err.Error())
	}
}

// Test if the client is able to get a specific registry
func TestGetRegistry(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	registries, err := c.GetRegistry("docker-registry")

	if err != nil || registries == nil {
		t.SkipNow()
	}
}

// Test if the client is able to create a registry
func TestCreateRegistry(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	username := "123456"
	password := "123456"

	err = c.CreateRegistry(
		CreateRegistryBody{
			AuthWithToken:       false,
			Filters:             []string{"*"},
			Name:                "docker-registry",
			Password:            &password,
			Registry:            "https://registry.hub.docker.com/",
			RegistryType:        "Docker Registry",
			RescanAfterDBUpdate: true,
			ScanLayers:          false,
			Username:            &username,
		},
	)

	if err != nil {
		t.Errorf(err.Error())
	}
}

// Test if the client is able to delete a specific registry
func TestDeleteRegistry(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	err = c.DeleteRegistry("unknown")

	if err != nil {
		t.SkipNow()
	}
}

// Test if the client is able to patch a registry
func TestPatchRegistry(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	username := ""
	password := ""

	err = c.PatchRegistry(
		CreateRegistryBody{
			AuthWithToken:       false,
			Filters:             []string{"neuvector/*"},
			Name:                "docker-registry",
			Password:            &password,
			Registry:            "https://registry.hub.docker.com/",
			RegistryType:        "Docker Registry",
			RescanAfterDBUpdate: false,
			ScanLayers:          true,
			Username:            &username,
		},
		"docker-registry",
	)

	if err != nil {
		t.SkipNow()
	}
}
