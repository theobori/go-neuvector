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
	auth := NewClientAuth(
		DefaultUsername,
		DefaultPassword,
	)
	
	return NewClient(
		&ClientConfig{
			auth: *auth,
			BaseUrl:  DefaultBaseUrl,
			Insecure: true,
		},
	)
}

// Used to set expired token with unitary tests
func (client *Client) SetToken(value string) {
	client.token = value
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

// Update the HTTP(s) request header with client configuration
func (client *Client) UpdateHeader(header *http.Header, hasBody bool) {
	header.Add("X-Auth-Token", client.token)
	header.Add("Accept", "application/json")

	if hasBody {
		header.Add("Content-Type", "application/json")
	}
}

// Create a basic NeuVector HTTP(s) request
func (client *Client) NewRequest(
	method string,
	endpoint string,
	reqBody any,
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

	client.UpdateHeader(&req.Header, body != nil)

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
// If there is no error, the JSON response is unmarshaled and stored into ret
func (client *Client) CallAPI(
	method string,
	endpoint string,
	reqBody any,
	ret any,
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

	if resp.StatusCode == http.StatusRequestTimeout {
		client.RefreshToken()

		return client.CallAPI(
			method,
			endpoint,
			reqBody,
			ret,
		)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return &APIError{
			resp.StatusCode,
			fmt.Sprintf("(%s) (%s)", method, endpoint),
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
func (client *Client) Get(endpoint string, ret any) error {
	return client.CallAPI("GET", endpoint, nil, ret)
}

// Computes a HTTP(s) POST request with the client context
func (client *Client) Post(endpoint string, reqBody any, ret any) error {
	return client.CallAPI("POST", endpoint, reqBody, ret)
}

// Computes a HTTP(s) PUT request with the client context
func (client *Client) Put(endpoint string, reqBody any, ret any) error {
	return client.CallAPI("PUT", endpoint, reqBody, ret)
}

// Computes a HTTP(s) DELETE request with the client context
func (client *Client) Delete(endpoint string, reqBody any, ret any) error {
	return client.CallAPI("DELETE", endpoint, reqBody, ret)
}

// Computes a HTTP(s) PATCH request with the client context
func (client *Client) Patch(endpoint string, reqBody any, ret any) error {
	return client.CallAPI("PATCH", endpoint, reqBody, ret)
}
