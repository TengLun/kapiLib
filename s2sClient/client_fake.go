package s2sclient

import (
	"net/http"
	"net/url"
)

// ClientFake struct to allow debugging of the S2SClient
type ClientFake struct {
	guid  string
	debug bool

	client  *http.Client
	baseURL *url.URL
}
