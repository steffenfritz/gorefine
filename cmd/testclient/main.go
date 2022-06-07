package main

import (
	"flag"
	"github.com/steffenfritz/gorefine"
	"log"
)

func main() {
	// For a full working client we should use a more sophisticated flag package.
	// However, for this test client we keep external dependencies manageable.
	serverurl := flag.String("s", "", "The URL of the OpenRefine server, e.g. http://localhost:8080")
	flag.Parse()

	if len(*serverurl) == 0 {
		log.Fatalln("-s flag is mandatory. Start this testclient with flag -h for help.")
	}

	// Create a new http client
	client := gorefine.NewClient(*serverurl)

	// Test GetProjectModel()
	// Expected output: risn
	pmodel, err := gorefine.GETProjectModel(client, "2525869207450")
	if err != nil {
		log.Println(err)
	} else {
		log.Println(pmodel.ColumnModel.Columns[0].Name)
	}
	// End Test GetProjetModel()

	// Test POSTExportRows()
	params := gorefine.ParamExportRows{
		ProjectID: "2525869207450",
		Format:    "csv",
	}

	form := gorefine.FormExportRows{
		Facets: []string{},
		Mode:   "row-based",
	}

	err = gorefine.POSTExportRows(client, params, form)
	if err != nil {
		log.Println(err.Error())
	}
	// End Test POSTExportRows()
}
