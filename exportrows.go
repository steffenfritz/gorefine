package gorefine

import (
	"encoding/json"
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
// and  'format' commonly csv, tsv, xls, xlsx, ods, html.
// In the form data it expects 'engine' : JSON string... (e.g. '{"facets":[],"mode":"row-based"}')
func POSTExportRows(ParamExportRows, FormExportRows) error {
	engine, err := json.Marshal(FormExportRows.Facets)
	if err != nil {
		return err
	}
	json.Valid([]byte(engine))

	return nil
}
