package neuvector

import (
	"testing"
)

// Test if the client is able to create a group
func TestCreateGroup(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	err = c.CreateGroup(
		CreateGroupBody{
			Name: "containerEQU",
			Criteria: []GroupCriteria{
				{
					Key:   "pattern",
					Value: "[0-9]",
					Op:    "regex",
				},
			},
			CfgType: "user_created",
		},
	)

	if err != nil {
		t.Errorf(err.Error())
	}
}

// Test if the client is able to get every group
func TestGetGroups(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	groups, err := c.GetGroups()

	if err != nil || groups == nil {
		t.Errorf(err.Error())
	}
}

// Test if the client is able to get a specifig group
func TestGetGroup(t *testing.T) {
	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	group, err := c.GetGroup("nodes")

	if err != nil || group == nil {
		t.Errorf(err.Error())
	}
}

// Test if the client is able to patch a specific a group
func TestPatchGroup(t *testing.T) {
	var err error

	c, err := NewDefaultClient()

	if err != nil {
		t.Errorf(err.Error())
		return
	}

	err = c.PatchGroup(
		"containerEQU",
		CreateGroupBody{
			Name: "containerEQU",
			Criteria: []GroupCriteria{
				{
					Key:   "pattern",
					Value: "[a-z]",
					Op:    "regex",
				},
			},
			CfgType: "user_created",
		},
	)

	if err != nil {
		t.Errorf(err.Error())
	}
}
