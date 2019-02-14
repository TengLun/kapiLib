// Package s2sclient is a package that allows sending user engagement information to
// the Kochava server. These transactions are broken down into either installs
// or events, which are effectively ations taken post-install. Events can have
// a substantial amount of meta-data that allows better analytics and user
// profile creation on the Kochava side.
//
// This package also allows sending of IdentityLink information, which is user
// metadata, and allows cross-device tracking.
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
func CreateClient(guid string, options ...func(*Client)) (APIAccessor, error) {

	var client Client

	client.guid = guid

	for _, option := range options {
		option(&client)
	}

	if client.debug == true {
		var cf ClientFake
		return cf, nil
	}

	if client.client == nil {
		client.client = &http.Client{}
	}

	return &client, nil
}
