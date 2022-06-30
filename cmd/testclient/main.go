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
	projectid := flag.String("id", "", "The project id value")
	flag.Parse()

	if len(*serverurl) == 0 {
		log.Fatalln("-s flag is mandatory. Start this testclient with flag -h for help.")
	}
	if len(*projectid) == 0 {
		log.Fatalln("-id flag is mandatory for this test client. Start this testclient with flag -h for help.")
	}

	// Create a new http client
	client := gorefine.NewClient(*serverurl)

	// Test GetProjectModel()
	// Expected output: risn
	pmodel, err := gorefine.GETProjectModel(client, *projectid)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(pmodel.ColumnModel.Columns[0].Name)
	}
	// End Test GetProjetModel()

	// Test POSTExportRows()
	params := gorefine.ParamExportRows{
		ProjectID: *projectid,
		Format:    "csv",
	}

	form := gorefine.FormExportRows{
		Facets: []string{},
		Mode:   "row-based",
	}

	// We always have to initialize a ParamTempl or ... use variadic args :(
	var templ gorefine.ParamTemplate
	// Test templated export
	templ.Separator = ","
	templ.Prefix = "BEGIN"
	templ.Suffix = "END"
	templ.TemplFile = "/Users/steffen/go/src/github.com/steffenfritz/gorefine/cmd/testclient/test.grel"

	// CSV export without template
	err = gorefine.POSTExportRows(client, params, form, templ)
	if err != nil {
		log.Println(err.Error())
	}
	// Template export
	params.Format = "template"
	err = gorefine.POSTExportRows(client, params, form, templ)
	if err != nil {
		log.Println(err.Error())
	}
	// End Test POSTExportRows()

	// Test GETCSRFToken()
	/*csrftoken, err := gorefine.GETCSRFToken(client)
	if err != nil {
		log.Fatalln(err.Error()) // We fatally quit here because we need the token for further tests
	}*/
	// Set generic parameters for altering requests
	/*var genericparams = gorefine.ParamGeneric{
		ProjectID: *projectid,
		CSRFToken: csrftoken.Token,
	}*/

	// Test POSTDeleteProject()
	/*result, err := gorefine.POSTDeleteProject(client, genericparams)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(result))*/
	// EndTest POSTDeleteProject
}
