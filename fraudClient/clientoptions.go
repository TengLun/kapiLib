package fraudclient

import (
	"net/http"
)

// SetHTTPClient allows client to be configured with custom httpClient
func SetHTTPClient(httpClient *http.Client) func(*Client) error {
	return func(client *Client) error {
		client.List.httpClient = httpClient
		client.Data.httpClient = httpClient
		return nil
	}
}

// SetDebugTrue returns a test client which returns spoof data instead of actually
// connecting to the API
func SetDebugTrue() func(*Client) error {
	return func(client *Client) error {
		client.debug = true
		client.Add.debug = true
		client.Remove.debug = true
		client.Update.debug = true
		client.List.debug = true
		client.Data.debug = true
		return nil
	}
}

// SetBaseURL allows client to be configured with custom httpClient
func SetBaseURL(BaseURL string) func(*Client) error {
	return func(client *Client) error {
		client.List.BaseURL = BaseURL
		client.Data.BaseURL = BaseURL
		return nil
	}
}
