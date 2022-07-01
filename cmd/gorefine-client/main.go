package main

import (
	flag "github.com/spf13/pflag"
	"github.com/steffenfritz/gorefine"
	"log"
)

func main() {

	var templ gorefine.ParamTemplate

	// Information flags
	serverurl := flag.StringP("server", "s", "", "The server URL, e.g. http://127.0.0.1:8080")
	projectid := flag.StringP("project-id", "", "", "The numeric project id")
	templfile := flag.StringP("template-file", "", "", "The path to a template file for exports")
	templatetext := flag.StringP("template", "", "", "The template provided as text argument for exports")
	templprefixfile := flag.StringP("prefix-file", "", "", "The path to a prefix template file for exports")
	templprefixtext := flag.StringP("prefix", "", "", "The template prefix provided as text argument for exports")
	templsuffixfile := flag.StringP("suffix-file", "", "", "The path to a suffix template file for exports")
	templsuffixtext := flag.StringP("suffix", "", "", "The template suffix provided as text argument for exports")
	format := flag.StringP("format", "f", "", "The format used for exports and imports, e.g. csv, tsv, json. If the format is template, templates can be used")
	facets := flag.StringSliceP("facets", "", []string{}, "A list of facets, e.g. --facets=\"v1,v2\"")
	mode := flag.StringP("mode", "", "row-based", "The data mode, possible values are row-based or record-based")
	// Action flags
	// getMetadata := flag.BoolP("get-metadata", "", false, "Get project metadata. project-id flag is mandatory")
	exportRows := flag.BoolP("export-rows", "", false, "Export rows. format flag is mandatory")
	// deleteProject := flag.BoolP("delete-project", "", false, "Delete Project. project-id flag is mandatory")

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

	// Set generic parameters for altering requests
	var genericparams = gorefine.ParamGeneric{
		ProjectID: *projectid,
		CSRFToken: csrftoken.Token,
	}

	// Export rows
	if *exportRows {
		if !(len(*format) == 0) {
			log.Println("The format flag is mandatory for this operation. Quitting.")
			return
		}

		if (*format == "template") && ((len(*templfile) == 0) || len(*templatetext) == 0) {
			log.Println("A template file or text is mandatory for this operation. Quittig.")
			return
		}
	}
}
