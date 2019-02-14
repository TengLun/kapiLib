package campaignclient

import (
	"net/http"
)

// SetHTTPClient allows client to be configured with custom httpClient
func SetHTTPClient(httpClient *http.Client) func(*Client) error {
	return func(client *Client) error {
		client.client = httpClient

		return nil
	}
}

// SetDebugTrue returns a test client which returns spoof data instead of actually
// connecting to the API
func SetDebugTrue() func(*Client) error {
	return func(client *Client) error {
		client.debug = true

		return nil
	}
}
