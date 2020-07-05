package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	c "github.com/damilarelana/goCYOA/core"
	m "github.com/damilarelana/goCYOA/middleware"
	"github.com/pkg/errors"
)

// define flags
var storyFilename *string = flag.String("file", "../storyData.json", "a json file containing Story chapters, arcs and options")
var templateFilename *string = flag.String("template", "../templates/storyChapter.gohtml", "a goHtml file used to render the json file data")
var port *int = flag.Int("port", 8085, "a port where the local webserver listens")

// defines the error message handler
func errMsgHandler(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func convertTemplateToString(t *template.Template) string {
	var temp bytes.Buffer
	var emptyChapterStruct c.Chapter
	err := t.Execute(&temp, emptyChapterStruct)
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to render goHTML template to a string %s\n", err.Error()))
	}
	return temp.String()
}

func parseTemplate(story c.Story) *template.Template {
	t, err := template.ParseFiles(*templateFilename)
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to parse goHTML file %s\n", err.Error()))
	}
	return t
}

func renderToStdout(t *template.Template, story c.Story) {
	for _, s := range story {
		err := t.Execute(os.Stdout, s)
		if err != nil {
			errMsgHandler(fmt.Sprintf("Failed to render goHTML file %s\n", err.Error()))
		}
	}
}

// define main function that:
//   * uses defaultMux()
//	 * initializes the flags
//   * parses the template
//   * renders the template
func main() {
	flag.Parse() // required to initialize the specified flags with the Operating system

	var story c.Story                              // initialize the `story	` struct i.e. without any data
	m.JSONFileHandler(storyFilename, &story)       // pass the initialized story struct and json-data storyFilename
	t := parseTemplate(story)                      // parseTemplate
	templateAsString := convertTemplateToString(t) // convert goHtml template to string
	fmt.Println(templateAsString)
	// renderToStdout(t, story) // render goHTML template to Stdout

	mux := m.DefaultMux() // create an instance of defaultMux()

	fmt.Println("\n==== ==== ==== ====")
	serverAddress := fmt.Sprintf("127.0.0.1:%d", *port)
	fmt.Println(fmt.Sprintf("Starting the webserver at http://%s\n", serverAddress))
	log.Fatal(errors.Wrap(http.ListenAndServe(serverAddress, mux), "Failed to start webserver"))
	// fmt.Printf("%+v\n", story)
}
