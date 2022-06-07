package gorefine

import (
	"net/http"
	"time"
)

const corepath = "/command/core/"

// Client is the client type used for every request and returned by NewClient
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient returns a Client struct
// webroot needs the schema, host and port, e.g. http://openrefine:3333
func NewClient(webroot string) *Client {
	return &Client{
		BaseURL: webroot,
		HTTPClient: &http.Client{
			Timeout: time.Minute, // timout set to one minute
		},
	}
}
