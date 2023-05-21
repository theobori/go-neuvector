package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/theobori/go-neuvector/logger"
)

const (
	// Default HTTP Content type
	DefaultContentType = "application/json"
	// Default REST API base url
	DefaultBaseUrl = "https://localhost:10443/v1"
	// Default Basic Auth username
	DefaultUsername = "admin"
	// Default Basic Auth password
	DefaultPassword = "admin"
)

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

// Create a basic NeuVector HTTP(s) request
func (client *Client) NewRequest(
	method string,
	endpoint string,
	reqBody interface{},
) (*http.Request, error) {
	var body []byte
	var err error

	url := client.BaseUrl + "/" + endpoint

	if reqBody != nil {
		body, err = json.Marshal(reqBody)

		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(
		method,
		url,
		bytes.NewReader(body),
	)

	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Auth-Token", client.token)
	req.Header.Add("Accept", "application/json")

	if body != nil {
		req.Header.Add("Content-Type", "application/json")
	}

	return req, nil
}

// Executes the HTTP(s) request then returns its response
func (client *Client) CallRequest(req *http.Request) (*http.Response, error) {
	resp, err := client.client.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Create a new NeuVector HTTP(s) request then executes it
//
// If there is no error, the reponse JSON is unmarshaled and stored into ret
func (client *Client) CallAPI(
	method string,
	endpoint string,
	reqBody interface{},
	ret interface{},
) error {
	req, err := client.NewRequest(
		method,
		endpoint,
		reqBody,
	)

	if err != nil {
		return err
	}

	resp, err := client.CallRequest(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return &APIError{
			resp.StatusCode,
			fmt.Sprintf("Unable to process (%s) (%s)", method, endpoint),
		}
	}

	if ret == nil {
		return nil
	}

	if err = json.Unmarshal([]byte(body), &ret); err != nil {
		return err
	}

	return nil
}

// Computes a HTTP(s) GET request with the client context
func (client *Client) Get(endpoint string, ret interface{}) error {
	return client.CallAPI("GET", endpoint, nil, ret)
}

// Computes a HTTP(s) POST request with the client context
func (client *Client) Post(endpoint string, reqBody interface{}, ret interface{}) error {
	return client.CallAPI("POST", endpoint, reqBody, ret)
}

// Computes a HTTP(s) PUT request with the client context
func (client *Client) Put(endpoint string, reqBody interface{}, ret interface{}) error {
	return client.CallAPI("PUT", endpoint, reqBody, ret)
}

// Computes a HTTP(s) DELETE request with the client context
func (client *Client) Delete(endpoint string, reqBody interface{}, ret interface{}) error {
	return client.CallAPI("DELETE", endpoint, reqBody, ret)
}

// Computes a HTTP(s) PATCH request with the client context
func (client *Client) Patch(endpoint string, reqBody interface{}, ret interface{}) error {
	return client.CallAPI("PATCH", endpoint, reqBody, ret)
}