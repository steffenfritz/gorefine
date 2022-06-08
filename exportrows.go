package gorefine

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

// ParamExportRows is the type for POSTExportRows parameters
type ParamExportRows struct {
	ProjectID string
	Format    string
}

// FormExportRows is the type for the form data used in POSTExportRows
type FormExportRows struct {
	Facets []string
	Mode   string
}

// POSTExportRows expects two parameters 'project', the project id
// and 'format' commonly csv, tsv, xls, xlsx, ods, html.
// In the form data it expects 'engine' : JSON string... (e.g. '{"facets":[],"mode":"row-based"}')
func POSTExportRows(c *Client, params ParamExportRows, forms FormExportRows) error {
	engine, err := json.Marshal(forms)
	if err != nil {
		return err
	}
	if jsonValid := json.Valid(engine); !jsonValid {
		return errors.New("JSON for export in POSTExportRows function not valid")
	}

	formData := url.Values{
		"project": {params.ProjectID},
		"format":  {params.Format},
		"engine":  {string(engine)},
	}

	resp, err := c.HTTPClient.PostForm(c.BaseURL+corepath+"export-rows", formData)
	if err != nil {
		return errors.New(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("HTTP Status code not ok: " + err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(err.Error())
	}

	// debug
	println(string(body))

	return nil
}
