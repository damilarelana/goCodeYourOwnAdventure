package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	c "github.com/damilarelana/goCYOA/core"
	m "github.com/damilarelana/goCYOA/middleware"
	"github.com/pkg/errors"
)

// define flags
var storyFilename *string = flag.String("file", "../storyData.json", "a json file containing Story chapters, arcs and options")
var templateFilename *string = flag.String("template", "../templates/gohtml/storyChapter.gohtml", "a goHtml file used to render the json file data")
var switchToWebServer *bool = flag.Bool("switch", false, "a flag that switches template render from `Stdout` to `webserver`. Default is `false` i.e. `Stdout`. To switch to webserver, use `-switch=true`")
var port *int = flag.Int("port", 8085, "a port where the local webserver listens")

// defines the error message handler
func errMsgHandler(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

/* customPathFunction() ...
- to be used with the Functional Option `WithCustomPathFn` in httpHandlers.go
- being used to show how to make the path different
- we use `/story/...` here instead of `/...`
*/
func customPathFunction(r *http.Request) (path string) {
	path = strings.TrimSpace(r.URL.Path)       // extract url path
	if path == "/story" || path == "/story/" { // ensures that root path always starts at the first chapter
		path = "/story/intro"
	}
	return path[len("/story/"):] // i.e. we need just `intro` from `/story/intro`
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
				<link rel="stylesheet" type="text/css" href="/templates/css/storyChapter.css">
			</head>
			<body>
				<section class="page">
					<h1>{{.Title}}</h1>
					{{range .Paragraph}} <!-- ranges over the story list -->
						<p>{{.}}</p> <!-- dumps all the data in that list element -->
					{{end}}
					<ul>
						{{range .Option}} <!-- range over the data in options -->
							<li><a href="/{{.Arc}}">{{.Text}}</a></li>
						{{end}}
					</ul>
				</section>
				<style>
					body {
						font-family: Arial, Helvetica, sans-serif;
					}
					h1 {
						text-align: center;
						position: relative;
					}
					.page {
						width: 80%;
						max-width: 500px;
						margin: auto;
						margin-top: 400px;
						margin-bottom: 40px;
						padding: 80px;
						background: #FFFCF6;
						border: 1px solid #eee;
						box-shadow: 0 10px 6px -6px #777;
					}
					ul{
						border-top: 1px dotted #ccc;
						padding-top: 10px 0 0 0;
						-webkit-padding-start: 0;
					}
					li {
						padding-top: 10px;
					}
					a,
					a:visited {
						text-decoration: none;
						color: #6295b5;
					}
					a:active,
					a:hover {
						color: #7792a2;
					}
					p {
						text-indent: 1em;
					}
			</style>
			</body>
		</html>`

		var storyTemplateAsString = `
		<!DOCTYPE html>
			<head>
				<meta charset="utf-8">
				<meta http-equiv="X-UA-Compatible" content="IE=edge">
				<title>Dynamic Adventure</title>
				<meta name="description" content="">
				<meta name="viewport" content="width=device-width, initial-scale=1">
				<link rel="stylesheet" type="text/css" href="/templates/css/storyChapter.css">
			</head>
			<body>
				<section class="page">
					<h1>{{.Title}}</h1>
					{{range .Paragraph}} <!-- ranges over the story list -->
						<p>{{.}}</p> <!-- dumps all the data in that list element -->
					{{end}}
					<ul>
						{{range .Option}} <!-- range over the data in options -->
							<li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
						{{end}}
					</ul>
				</section>
				<style>
					body {
						font-family: Arial, Helvetica, sans-serif;
					}
					h1 {
						text-align: center;
						position: relative;
					}
					.page {
						width: 80%;
						max-width: 500px;
						margin: auto;
						margin-top: 400px;
						margin-bottom: 40px;
						padding: 80px;
						background: #FFFCF6;
						border: 1px solid #eee;
						box-shadow: 0 10px 6px -6px #777;
					}
					ul{
						border-top: 1px dotted #ccc;
						padding-top: 10px 0 0 0;
						-webkit-padding-start: 0;
					}
					li {
						padding-top: 10px;
					}
					a,
					a:visited {
						text-decoration: none;
						color: #6295b5;
					}
					a:active,
					a:hover {
						color: #7792a2;
					}
					p {
						text-indent: 1em;
					}
			</style>
			</body>
		</html>`

		/* create an instance of CustomHandler()
		- we can use `m.WithTemplate(InitTemplateForWeb(templateAsString))` to pass in another template here
		- by passing it into the `CustomHandler` as the `opts`
			+	mux := m.CustomHandler(&story, templateAsString, m.WithTemplate(InitTemplateForWeb(templateAsString))
			+ instead of just `mux := m.CustomHandler(&story, templateAsString)`
		*/
		customMux := m.CustomHandler(&story, templateAsString, m.WithTemplate(m.InitTemplateForWeb(storyTemplateAsString)), m.WithCustomPathFn(customPathFunction))

		/*
			Initialize the DefaultMux [to help with 404 page handling]
			passing the
		*/
		mux := m.DefaultMux(customMux)

		fmt.Println("\n==== ==== ==== ====")
		serverAddress := fmt.Sprintf("127.0.0.1:%d", *port)
		fmt.Println(fmt.Sprintf("\nStarting the webserver at http://%s\n", serverAddress))
		log.Fatal(errors.Wrap(http.ListenAndServe(serverAddress, mux), "Failed to start webserver"))
		// fmt.Printf("%+v\n", story)
	} else {
		m.RenderToStdout(t, story) // render goHTML template to Stdout
	}
}
