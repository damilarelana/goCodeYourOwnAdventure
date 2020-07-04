package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"

	c "github.com/damilarelana/goCYOA/core"
	m "github.com/damilarelana/goCYOA/middleware"
)

// define flags
var storyFilename *string = flag.String("file", "../storyData.json", "a json file containing Story chapters, arcs and options")

// defines the error message handler
func errMsgHandler(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func renderTemplate(story c.Story) {
	t, err := template.ParseFiles("../templates/storyChapter.gohtml")
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to parse goHTML file %s\n", err.Error()))
	}
	for _, s := range story {
		err = t.Execute(os.Stdout, s)
		if err != nil {
			errMsgHandler(fmt.Sprintf("Failed to render goHTML file %s\n", err.Error()))
		}
	}
}

func main() {
	// parse flags
	flag.Parse()                             // required to initialize the specified flags with the Operating system
	var story c.Story                        // initialize the `story	` struct
	m.JSONFileHandler(storyFilename, &story) // pass the initialized story struct and json-data storyFilename
	renderTemplate(story)                    // render goHTML template
	// fmt.Printf("%+v\n", story)
}
