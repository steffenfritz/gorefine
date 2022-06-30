package gorefine

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// ParamExportRows is the type for POSTExportRows parameters
type ParamExportRows struct {
	ProjectID string
	Format    string
}

// ParamTemplate holds the parameters for templated exports
type ParamTemplate struct {
	Prefix    string
	Suffix    string
	Separator string
	TemplFile string
}

// FormExportRows is the type for the form data used in POSTExportRows
type FormExportRows struct {
	Facets []string
	Mode   string
}

// POSTExportRows expects two parameters 'project', the project id
// and 'format' commonly csv, tsv, xls, xlsx, ods, html.
// In the form data it expects 'engine' : JSON string... (e.g. '{"facets":[],"mode":"row-based"}')
func POSTExportRows(c *Client, params ParamExportRows, forms FormExportRows, templ ParamTemplate) error {
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

	if params.Format == "template" {
		// We read the template file from the provided path
		// ioutil.ReadFile should be ok for template files
		content, err := ioutil.ReadFile(templ.TemplFile)
		if err != nil {
			return err
		}
		formData.Add("prefix", templ.Prefix)
		formData.Add("suffix", templ.Suffix)
		formData.Add("separator", templ.Separator)
		formData.Add("template", string(content))
	}

	resp, err := c.HTTPClient.PostForm(c.BaseURL+corepath+"export-rows", formData)
	if err != nil {
		return errors.New(err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("HTTP Status code not ok: " + strconv.Itoa(resp.StatusCode))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.New(err.Error())
	}

	// debug
	println(string(body))

	return nil
}
