package main

import (
	"flag"
	"fmt"
	"os"

	c "github.com/damilarelana/goCYOA"
	m "github.com/damilarelana/goCYOA/middleware"
)

// define flags
var storyFilename *string = flag.String("file", "../../storyData.json", "a json file containing Story chapters, arcs and options")

// defines the error message handler
func errMsgHandler(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {
	// parse flags
	flag.Parse()                            // required to initialize the specified flags with the Operating system
	var story c.Story                       // initialize the `story	` struct
	m.JSONFileHandler(storyFilename, &story) // pass the initialized story struct and json-data storyFilename

	// print content
	fmt.Printf("%+v\n", story)
}
