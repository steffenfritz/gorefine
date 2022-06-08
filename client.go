package gorefine

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

const corepath = "/command/core/"

// Client is the client type used for every request and returned by NewClient
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// csrftoken type
type csrftoken struct {
	Token string `json:"token"`
}

// ParamGeneric is a type holding project id and csrf token
// as this are used all over the place
type ParamGeneric struct {
	ProjectID string
	CSRFToken string
}

// GenericFormData returns a form filled with generic information
// project : projectid. csrf_token : CSRF token
func GenericFormData(params ParamGeneric) url.Values {
	return url.Values{
		"project":    {params.ProjectID},
		"csrf_token": {params.CSRFToken},
	}
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

// GETCSRFToken returns a CSRF token that can be used with
// POST requests that alter the state of the OpenRefine server
func GETCSRFToken(c *Client) (csrftoken, error) {
	var csrftoken csrftoken

	resp, err := c.HTTPClient.Get(c.BaseURL + corepath + "get-csrf-token")
	if err != nil {
		return csrftoken, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return csrftoken, err
	}
	err = json.Unmarshal(body, &csrftoken)

	return csrftoken, err
}
