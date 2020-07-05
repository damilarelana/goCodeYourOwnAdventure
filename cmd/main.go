package main

import (
	"flag"
	"fmt"
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

// define main function that:
//   * uses defaultMux()
//	 * initializes the flags
//   * parses the template
//   * renders the template
func main() {
	flag.Parse() // required to initialize the specified flags with the Operating system

	var story c.Story                                // initialize the `story	` struct i.e. without any data
	m.JSONFileHandler(storyFilename, &story)         // pass the initialized story struct and json-data storyFilename
	t := m.ParseTemplate(story, templateFilename)    // parseTemplate
	templateAsString := m.ConvertTemplateToString(t) // convert goHtml template to string
	fmt.Println(templateAsString)
	m.RenderToStdout(t, story) // render goHTML template to Stdout

	mux := m.DefaultMux() // create an instance of defaultMux()

	fmt.Println("\n==== ==== ==== ====")
	serverAddress := fmt.Sprintf("127.0.0.1:%d", *port)
	fmt.Println(fmt.Sprintf("Starting the webserver at http://%s\n", serverAddress))
	log.Fatal(errors.Wrap(http.ListenAndServe(serverAddress, mux), "Failed to start webserver"))
	// fmt.Printf("%+v\n", story)
}
