package s2sclient

import (
	"net/http"
	"net/url"
)

// Client primary struct to access all of the available methods
type Client struct {
	guid string

	httpClient *http.Client
	baseURL    *url.URL
}

// CreateClient creates the client to use the methods in the API
func CreateClient(guid string, options ...func(*Client)) (*Client, error) {

	var client Client

	client.guid = guid

	for _, option := range options {
		option(&client)
	}

	if client.httpClient == nil {
		client.httpClient = &http.Client{}
	}

	return &client, nil
}
