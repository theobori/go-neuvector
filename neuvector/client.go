package neuvector

import (
	"log"
	"net/http"

	"github.com/theobori/go-neuvector/logger"
)

// Default HTTP Content type
const DefaultContentType = "application/json"

// Default REST API base url
const DefaultBaseUrl = "https://localhost:10443/v1"

// Default Basic Auth username
const DefaultUsername = "admin"

// Default Basic Auth password
const DefaultPassword = "admin"

// Client primary data to operate with NeuVector
type Client struct {
	// Controller REST API base url
	BaseUrl string
	// Logger
	Logger *log.Logger
	// HTTP configured client
	client *http.Client
	// Authentication token
	token string
	// Client configuration, only used to refresh the token
	config *ClientConfig
}

// Instanciates a new Client object
func NewClient(config *ClientConfig) (*Client, error) {
	client := Client{
		BaseUrl: config.BaseUrl,
		Logger:  logger.Info,
		token:   "",
		client:  config.GetHTTPClient(),
		config:  config,
	}

	if err := client.RefreshToken(); err != nil {
		return nil, err
	}

	return &client, nil
}

// Instanciates a new Client object with its default values
func NewDefaultClient() (*Client, error) {
	basicAuth := BasicAuth{
		Username: DefaultUsername,
		Password: DefaultPassword,
	}

	return NewClient(
		&ClientConfig{
			Auth: ClientAuth{
				basicAuth,
			},
			BaseUrl:  DefaultBaseUrl,
			insecure: true,
		},
	)
}

// Refresh the client token
func (client *Client) RefreshToken() error {
	tokenResponse, err := client.config.GetToken()

	if err != nil {
		return err
	}

	client.token = tokenResponse.Getvalue()

	return nil
}
