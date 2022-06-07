package gorefine

// ProjectMD is the exported project metadata type
type ProjectMD struct {
	ProjectId struct {
		Name           string `json:"name"`
		Created        string `json:"created"`
		Modified       string `json:"modified"`
		CustomMetadata struct {
		} `json:"customMetadata"`
	}
}

// ProjectsMD is the type returned by GetProjectMD
// called via /command/core/get-all-project-metadata
type ProjectsMD struct {
	Projects []ProjectsMD
}

// GetProjectMD is a GET to request all project metadata.
// This request has no parameters.
func GetProjectMD(c *Client) {}
