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

func (c *Client) GetUser(fullname string) (*User, error) {
	var ret GetUserResponseFull

	url := fmt.Sprintf("/user/%s", fullname)

	if err := c.Get(url, &ret); err != nil {
		return nil, err
	}

	return &ret.User, nil
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