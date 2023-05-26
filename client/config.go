package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
)

// NeuVector Controller credentials
type BasicAuth struct {
	// Client username
	Username string `json:"username"`
	// Client password
	Password string `json:"password"`
}

// Needed JSON representation to authenticate
type ClientAuth struct {
	// JSON autentication object
	Password BasicAuth `json:"password"`
}

// Returns *ClientAuth
func NewClientAuth(username string, password string) *ClientAuth {
	return &ClientAuth{
		Password: BasicAuth{
			Username: username,
			Password: password,
		},
	}
}

// Client configuration
type ClientConfig struct {
	// TLS skip invalid certificate
	Insecure bool
	// Controller REST API base url
	BaseUrl string
	// Crednetials
	auth ClientAuth
}

// Returns *ClientConfig
func NewClientConfig(auth *ClientAuth) *ClientConfig {
	return &ClientConfig{
		auth: *auth,
	}
}

// Represents the NeuVector token JSON object
//
// Avoiding every other fields, we only target the token one
type TokenResponse struct {
	// Token root field
	Token struct {
		// Token value
		Token string `json:"token"`
	} `json:"token"`
}

// Returns the token value
func (tokenResponse *TokenResponse) Getvalue() string {
	return tokenResponse.Token.Token
}

// Returns the transport configuration for the HTTP client
func (config *ClientConfig) GetHTTPTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.Insecure,
		},
	}
}

// Returns the HTTP client used to send requests
func (config *ClientConfig) GetHTTPClient() *http.Client {
	return &http.Client{
		Transport: config.GetHTTPTransport(),
	}
}

// Get the client token required to interact with the Controller REST API
func (config *ClientConfig) GetToken() (TokenResponse, error) {
	var token TokenResponse

	client := config.GetHTTPClient()
	url := config.BaseUrl + "/auth"
	bufferJSON, err := json.Marshal(config.auth)

	if err != nil {
		return token, err
	}

	reader := bytes.NewReader(bufferJSON)
	resp, err := client.Post(
		url,
		DefaultContentType,
		reader,
	)

	if err != nil {
		return token, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return token, &APIError{
			resp.StatusCode,
			"Invalid authentication",
		}
	}

	// Get the token in the JSON body
	body, err := io.ReadAll(resp.Body)

	// Unmarshal or Decode the JSON into a TokenResponse
	json.Unmarshal([]byte(body), &token)

	return token, err
}
