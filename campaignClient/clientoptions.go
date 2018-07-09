package campaignclient

import (
	"net/http"
)

// SetHTTPClient allows client to be configured with custom httpClient
func SetHTTPClient(httpClient *http.Client) func(*aPIA) error {
	return func(client *aPIA) error {
		client.client = httpClient
		return nil
	}
}

// SetDebugTrue returns a test client which returns spoof data instead of actually
// connecting to the API
func SetDebugTrue() func(*aPIA) error {
	return func(client *aPIA) error {
		client.debug = true
		return nil
	}
}
