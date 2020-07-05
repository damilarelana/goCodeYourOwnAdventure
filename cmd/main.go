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
var switchToWebServer *bool = flag.Bool("switch", false, "a flag that switches template render from `Stdout` to `webserver`. Default is `false` i.e. `Stdout`. To switch to webserver, use `-switch=true`")
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

	var story c.Story                        // initialize the `story	` struct i.e. without any data
	m.JSONFileHandler(storyFilename, &story) // parse the json-data (i.e. storyFilename) into the initialized story struct)
	t := m.ParseTemplate(templateFilename)   // parse the template

	if *switchToWebServer { // checks if flag says we should switch to webServer i.e. true, then initializes code required to rendere to webserver
		// templateAsString := m.ConvertTemplateToString(t) // convert the parsed template to a string
		// fmt.Println("\n==== Template ====")
		// fmt.Println(templateAsString)

		var templateAsString = `
			<!DOCTYPE html>
			<head>
				<meta charset="utf-8">
				<meta http-equiv="X-UA-Compatible" content="IE=edge">
				<title>Dynamic Adventure</title>
				<meta name="description" content="">
				<meta name="viewport" content="width=device-width, initial-scale=1">
			</head>
			<body>
				<h1>{{.Title}}</h1>
				{{range .Paragraph}} <!-- ranges over the story list -->
					<p>{{.}}</p> <!-- dumps all the data in that list element -->
				{{end}}
				<ul>
					{{range .Option}} <!-- range over the data in options -->
						<li><a href="/{{.Arc}}">{{.Text}}</a></li>
					{{end}}
				</ul>
			</body>
		</html>`

		mux := m.CustomHandler(&story, templateAsString) // create an instance of CustomHandler()

		fmt.Println("\n==== ==== ==== ====")
		serverAddress := fmt.Sprintf("127.0.0.1:%d", *port)
		fmt.Println(fmt.Sprintf("\nStarting the webserver at http://%s\n", serverAddress))
		log.Fatal(errors.Wrap(http.ListenAndServe(serverAddress, mux), "Failed to start webserver"))
		// fmt.Printf("%+v\n", story)
	} else {
		m.RenderToStdout(t, story) // render goHTML template to Stdout
	}
}
