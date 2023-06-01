package neuvector

import (
	"fmt"
	"testing"
)

// Test if the client is able to get a specific user
func TestGetUser(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	user, err := c.GetUser("admin")

	if err != nil || user == nil {
		t.Error(err)
	}
}

// Test if the client is able to create an user
func TestCreateUser(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	password := "SuperPassword1&"

	err = c.CreateUser(
		CreateUserBody{
			Username: "usertest",
			Fullname: "usertest",
			Email:    "my-email@free.fr",
			Locale:   "en",
			Password: &password,
			Role:     "reader",
			Timeout:  300,
		},
	)

	if err != nil {
		t.Error(err)
	}
}

// Test if the client is able to patch a specific user
func TestPatchUser(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	role := "admin"

	err = c.PatchUser(
		"usertest",
		PatchUserBody{
			Fullname: "usertest",
			Role:     &role,
		},
	)

	if err != nil {
		t.Error(err)
	}
}

// Test if the client is able to delete a specific user
func TestDeleteUser(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	err = c.DeleteUser("usertest")

	if err != nil {
		t.Error(err)
	}
}

// Test if the client is able to get a specific user role
func TestGetUserRole(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	role, err := c.GetUserRole("reader")

	fmt.Println(role)

	if err != nil || role == nil {
		t.Error(err)
	}
}

// Test if the client is able to create an user role
func TestCreateUserRole(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	err = c.CreateUserRole(
		CreateUserRoleBody{
			Name:    "customtest",
			Comment: "",
			Permissions: []UserRolePermission{
				{
					ID:    "vulnerability",
					Read:  true,
					Write: true,
				},
			},
		},
	)

	if err != nil {
		t.Error(err)
	}
}

// Test if the client is able to create an user role
func TestPatchUserRole(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	err = c.PatchUserRole(
		"customtest",
		PatchUserRoleBody{
			Name:    "customtest",
			Comment: "HHHHHHHHHH",
			Permissions: []UserRolePermission{
				{
					ID:    "ci_scan",
					Read:  false,
					Write: true,
				},
			},
		},
	)

	if err != nil {
		t.Error(err)
	}
}

// Test if the client is able to delete a specific user role
func TestDeleteUserRole(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Error(err)
		return
	}

	err = c.DeleteUserRole("customtest")

	if err != nil {
		t.Error(err)
	}
}
