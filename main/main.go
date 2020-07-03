package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	cyoa "github.com/damilarelana/goCYOA"
)

// define flags
var storyFilename *string = flag.String("file", "storyData.json", "a json file containing Story chapters, arcs and options")

// defines the error message handler
func errMsgHandler(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// openFile() uses the system to open and prep the file for reading
func openFile(storyFilename *string) *os.File {
	openedFile, err := os.Open(*storyFilename)
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to open JSON file: %s\n", *storyFilename))
	}
	return openedFile
}

// readFile()
//  * takes the pointer to the opened file i.e. openedFile
func readFile(f *os.File) *json.Decoder {
	fileData := json.NewDecoder(f)
	return fileData
}

// parseFileData parses the jsonData into structData
func parseFileData(fileData *json.Decoder, story *cyoa.Story) {
	err := fileData.Decode(&story) // decode and parse the `json data` into `structs data` within the address space &story
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to parse file data %v\n", *fileData))
	}
}

func main() {
	// parse flags
	flag.Parse()                          // required to initialize the specified flags with the Operating system
	var story cyoa.Story                  // initialize the `story	` struct
	openedFile := openFile(storyFilename) // open file
	fileData := readFile(openedFile)      // read file content
	parseFileData(fileData, &story)       // parse file data

	// print content
	fmt.Printf("%+v\n", story)
}
