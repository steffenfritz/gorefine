package gorefine

import (
	"encoding/json"
)

// ProjectModel is the type that holds the model for a specific project
// This includes columns, records, overlay models and scripting
type ProjectModel struct {
	ColumnModel struct {
		Columns []struct {
			CellIndex    int    `json:"cellIndex"`
			OriginalName string `json:"originalName"`
			Name         string `json:"name"`
		} `json:"columns"`
		KeyCellIndex  int           `json:"keyCellIndex"`
		KeyColumnName string        `json:"keyColumnName"`
		ColumnGroups  []interface{} `json:"columnGroups"`
	} `json:"columnModel"`
	RecordModel struct {
		HasRecords bool `json:"hasRecords"`
	} `json:"recordModel"`
	OverlayModels struct {
	} `json:"overlayModels"`
	Scripting struct {
		Grel struct {
			Name              string `json:"name"`
			DefaultExpression string `json:"defaultExpression"`
		} `json:"grel"`
		Jython struct {
			Name              string `json:"name"`
			DefaultExpression string `json:"defaultExpression"`
		} `json:"jython"`
		Clojure struct {
			Name              string `json:"name"`
			DefaultExpression string `json:"defaultExpression"`
		} `json:"clojure"`
	} `json:"scripting"`
}

// GETProjectModel expects a project id via query string in the url.
// It returns a ProjectModel and an error.
func GETProjectModel(c *Client, projectid string) (ProjectModel, error) {
	var pm ProjectModel
	resp, err := c.HTTPClient.Get(c.BaseURL + corepath + "get-models?project=" + projectid)
	if err != nil {
		return pm, err
	}
	defer resp.Body.Close()

	// Try to decode the response to the struct
	err = json.NewDecoder(resp.Body).Decode(&pm)
	if err != nil {
		return pm, err
	}

	return pm, nil
}
