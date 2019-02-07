package s2sclient

import (
	"net/http"
	"net/url"
)

// Client primary struct to access all of the available methods
type ClientFake struct {
	guid  string
	debug bool

	client  *http.Client
	baseURL *url.URL
}
