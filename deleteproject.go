package gorefine

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// POSTDeleteProject expects a project id in a json encoded body and return a json encoded response
func POSTDeleteProject(c *Client, params ParamGeneric) ([]byte, error) {
	formData := GenericFormData(params)

	resp, err := c.HTTPClient.PostForm(c.BaseURL+"/command/core/delete-project", formData)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, errors.New("HTTP Status code not ok: " + err.Error())

	}

	// As the response is not documented we just return it as an array of bytes. This is ugly but for god's sake...
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, err
}
