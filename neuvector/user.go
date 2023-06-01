package neuvector

import "fmt"

type RoleDomains struct {
	Role    string   `json:"role"`
	Domains []string `json:"domains"`
}

type PatchUserBody struct {
	Fullname    string       `json:"fullname"`
	Password    *string      `json:"password,omitempty"`
	NewPassword *string      `json:"new_password,omitempty"`
	PwdProfile  string       `json:"pwd_profile"`
	Email       *string      `json:"email,omitempty"`
	Role        *string      `json:"role,omitempty"`
	Timeout     *int         `json:"timeout,omitempty"`
	Locale      *string      `json:"locale,omitempty"`
	RoleDomains *RoleDomains `json:"role_domains,omitempty"`
}

type PatchUserBodyFull struct {
	Config PatchUserBody `json:"config"`
}

type getUserResponse struct {
	Fullname                  string       `json:"fullname"`
	Server                    string       `json:"server"`
	Username                  string       `json:"username"`
	Password                  *string      `json:"password,omitempty"`
	Email                     string       `json:"email"`
	Role                      string       `json:"role"`
	Timeout                   int          `json:"timeout"`
	Locale                    string       `json:"locale"`
	DefaultPassword           bool         `json:"default_password"`
	ModifyPassword            bool         `json:"modify_password"`
	RoleDomains               *RoleDomains `json:"role_domains,omitempty"`
	LastLoginTimestamp        int64        `json:"last_login_timestamp"`
	LastLoginAt               string       `json:"last_login_at"`
	LoginCount                uint32       `json:"login_count"`
	BlockedForFailedLogin     bool         `json:"blocked_for_failed_login"`
	BlockedForPasswordExpired bool         `json:"blocked_for_password_expired"`
}

type User = getUserResponse

type GetUserResponseFull struct {
	User User `json:"user"`
}

type CreateUserBody = User

type CreateUserBodyFull = GetUserResponseFull

type UserRolePermission struct {
	ID    string `json:"id"`
	Read  bool   `json:"read"`
	Write bool   `json:"write"`
}

type GetUserRoleResponse struct {
	Name        string               `json:"name"`
	Comment     string               `json:"comment"`
	Reserved    bool                 `json:"reserved"`
	Permissions []UserRolePermission `json:"permissions"`
}

type UserRole = GetUserRoleResponse

type GetUserRoleResponseFull struct {
	Role UserRole `json:"role"`
}

type CreateUserRoleBody struct {
	Name        string               `json:"name"`
	Comment     string               `json:"comment"`
	Permissions []UserRolePermission `json:"permissions"`
}

type CreateUserRoleBodyFull struct {
	Config CreateUserRoleBody `json:"config"`
}

type PatchUserRoleBody = CreateUserRoleBody

type PatchUserRoleBodyFull = CreateUserRoleBodyFull

func (c *Client) GetUser(fullname string) (*GetUserResponseFull, error) {
	var ret GetUserResponseFull

	url := fmt.Sprintf("/user/%s", fullname)

	if err := c.Get(url, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) CreateUser(body CreateUserBody) error {
	return c.Post(
		"/user",
		CreateUserBodyFull{body},
		nil,
	)
}

func (c *Client) DeleteUser(fullname string) error {
	return c.Delete(
		fmt.Sprintf("/user/%s", fullname),
		nil,
		nil,
	)
}

func (c *Client) PatchUser(fullname string, body PatchUserBody) error {
	return c.Patch(
		fmt.Sprintf("/user/%s", fullname),
		PatchUserBodyFull{body},
		nil,
	)
}

func (c *Client) GetUserRole(name string) (*GetUserRoleResponseFull, error) {
	var ret GetUserRoleResponseFull

	url := fmt.Sprintf("/user_role/%s", name)

	if err := c.Get(url, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

func (c *Client) CreateUserRole(body CreateUserRoleBody) error {
	return c.Post(
		"/user_role",
		CreateUserRoleBodyFull{body},
		nil,
	)
}

func (c *Client) DeleteUserRole(name string) error {
	return c.Delete(
		fmt.Sprintf("/user_role/%s", name),
		nil,
		nil,
	)
}

func (c *Client) PatchUserRole(fullname string, body PatchUserRoleBody) error {
	return c.Patch(
		fmt.Sprintf("/user_role/%s", fullname),
		PatchUserRoleBodyFull{body},
		nil,
	)
}
