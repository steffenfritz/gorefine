package main

import (
	"encoding/json"
	"fmt"
	"log"

	flag "github.com/spf13/pflag"
	"github.com/steffenfritz/gorefine"
)

func main() {

	var templ gorefine.ParamTemplate

	// Information flags
	serverurl := flag.StringP("server", "s", "", "The server URL, e.g. http://127.0.0.1:8080")
	projectid := flag.StringP("project-id", "", "", "The numeric project id")
	templ.TemplFile = flag.StringP("template-file", "", "", "The path to a template file for exports")
	templ.Template = flag.StringP("template", "", "", "The template provided as text argument for exports")
	templ.PrefixFile = flag.StringP("prefix-file", "", "", "The path to a prefix template file for exports")
	templ.Prefix = flag.StringP("prefix", "", "", "The template prefix provided as text argument for exports")
	templ.SuffixFile = flag.StringP("suffix-file", "", "", "The path to a suffix template file for exports")
	templ.Suffix = flag.StringP("suffix", "", "", "The template suffix provided as text argument for exports")
	templ.Separator = flag.StringP("separator", "", "", "A separator provided as text argument for exports")
	format := flag.StringP("format", "f", "", "The format used for exports and imports, e.g. csv, tsv, json. If the format is template, templates can be used")
	facets := flag.StringSliceP("facets", "", []string{}, "A list of facets, e.g. --facets=\"v1,v2\"")
	mode := flag.StringP("mode", "", "row-based", "The data mode, possible values are row-based or record-based")
	// Action flags
	//getMetadata := flag.BoolP("get-metadata", "", false, "Get project metadata. project-id flag is mandatory")
	getModel := flag.BoolP("get-model", "", false, "Get project model. project-id flag is mandatory")
	exportRows := flag.BoolP("export-rows", "", false, "Export rows. format flag is mandatory")
	deleteProject := flag.BoolP("delete-project", "", false, "Delete Project. project-id flag is mandatory")

	printLicense := flag.Bool("license", false, "Print the license")
	flag.Parse()

	if *printLicense {
		printlicense()
		return
	}

	if len(*serverurl) == 0 {
		log.Println("The server flag is mandatory. Quitting.")
		return
	}

	// Create a new http client
	client := gorefine.NewClient(*serverurl)

	// GETCSRFToken()
	// The token is not needed for every request
	csrftoken, err := gorefine.GETCSRFToken(client)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Set static parameters for most of the requests
	var genericparams = gorefine.ParamGeneric{
		ProjectID: *projectid,
		CSRFToken: csrftoken.Token,
	}

	// Export rows function
	if *exportRows {
		if len(*format) == 0 {
			log.Println("The format flag is mandatory for this operation. Quitting.")
			return
		}

		if (*format == "template") && ((len(*templ.TemplFile) == 0) && len(*templ.Template) == 0) {
			log.Println("A template file or text is mandatory for this operation. Quitting.")
			return
		}
		params := gorefine.ParamExportRows{
			ProjectID: *projectid,
			Format:    *format,
		}
		form := gorefine.FormExportRows{
			Facets: *facets,
			Mode:   *mode,
		}
		err = gorefine.POSTExportRows(client, params, form, templ)
		if err != nil {
			log.Println(err.Error())
		}
	}

	// Get project model
	if *getModel {
		model, err := gorefine.GETProjectModel(client, *projectid)
		if err != nil {
			log.Println(err.Error())
		}
		modelJSON, err := json.Marshal(model)
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Println(string(modelJSON))
	}

	// Delete project
	if *deleteProject {
		result, err := gorefine.POSTDeleteProject(client, genericparams)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(string(result))
	}

}
