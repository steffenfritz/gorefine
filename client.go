package gorefine

import (
	"net/http"
	"time"
)

// Client is the client type used for every request and returned by NewClient
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient returns a Client struct
// webroot needs the schema, host and port, e.g. http://openrefine:3333
func NewClient(webroot string) *Client {
	baseurl := webroot + "/command/core/"
	return &Client{
		BaseURL: baseurl,
		HTTPClient: &http.Client{
			Timeout: time.Minute, // timout set to one minute
		},
	}
}
