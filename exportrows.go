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
	Prefix     *string
	Suffix     *string
	Template   *string
	Separator  *string
	TemplFile  *string
	PrefixFile *string
	SuffixFile *string
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

	// We have to decide some steps here:
	//   If we have the template format, we have to check if there are prefix and suffix files
	//   If so we ignore prefixes and suffixes provided via cli argument.
	if params.Format == "template" {
		var templtemplate string
		var templprefix string
		var templsuffix string

		if len(*templ.TemplFile) != 0 {
			// We read the template files from the provided paths
			// ioutil.ReadFile should be ok for template files
			// If no template file is provided we set the template string to templ.Template
			// If no templ.Template is provided there should be no problem as it is initialized
			// as an empty string.
			tmpTempltemplate, err := ioutil.ReadFile(*templ.TemplFile)
			if err != nil {
				return errors.New("Could not read template file: " + err.Error())
			}
			templtemplate = string(tmpTempltemplate)
		} else {
			templtemplate = *templ.Template
		}

		if len(*templ.PrefixFile) != 0 {
			tmpTemplprefix, err := ioutil.ReadFile(*templ.PrefixFile)
			if err != nil {
				return errors.New("Could not read prefix file: " + err.Error())
			}
			templprefix = string(tmpTemplprefix)
		} else {
			templprefix = *templ.Prefix
		}
		if len(*templ.SuffixFile) != 0 {
			tmpTemplsuffix, err := ioutil.ReadFile(*templ.SuffixFile)
			if err != nil {
				return errors.New("Could not read suffix file: " + err.Error())
			}
			templsuffix = string(tmpTemplsuffix)
		} else {
			// We handle the suffix like the prefix file
			templsuffix = *templ.Suffix
		}
		if err != nil {
			return err
		}
		formData.Add("prefix", templprefix)
		formData.Add("suffix", templsuffix)
		formData.Add("separator", *templ.Separator)
		formData.Add("template", string(templtemplate))
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
