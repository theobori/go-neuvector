package neuvector

import (
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

	err = c.CreateUser(
		CreateUserBody{
			Username: "usertest",
			Fullname: "usertest",
			Email: "my-email@free.fr",
			Locale: "en",
			Password: "SuperPassword1&",
			Role: "reader",
			Timeout: 300,
		},
	)

	if err != nil{
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