package gorefine

import (
	"bytes"
	"io/ioutil"
)

// POSTDeleteProject expects a project id in a json encoded body and return a json encoded response
func POSTDeleteProject(c *Client, projectid string) ([]byte, error) {
	jsonData := []byte(`{"project":"` + projectid + `"}`)
	resp, err := c.HTTPClient.Post(c.BaseURL+"/command/core/delete-project", "application/json", bytes.NewBuffer(jsonData))
	defer resp.Body.Close()

	// As the response is not documented we just return it as an array of bytes. This is ugly but for god's sake...
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, err
}
