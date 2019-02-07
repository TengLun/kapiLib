package s2sclient

import (
	"net/http"
	"net/url"
)

// Client primary struct to access all of the available methods
type Client struct {
	guid  string
	debug bool

	client  *http.Client
	baseURL *url.URL
}

// CreateClient creates the client to use the methods in the API
func CreateClient(guid string, options ...func(*Client)) (*Client, error) {

	var client Client

	client.guid = guid

	for _, option := range options {
		option(&client)
	}

	if client.client == nil {
		client.client = &http.Client{}
	}

	return &client, nil
}
