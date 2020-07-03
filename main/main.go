package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

// define flags
var storyFilename *string = flag.String("file", "storyData.json", "a json file containing Story chapters, arcs and options")

// defines the error message handler
func errMsgHandler(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// define parser to
func parseFileData(fileData *json.Decoder) []problem {
	returnedValue := make([]problem, len(records))
	for i, record := range records { // iterate over the multi-dimensional slice
		returnedValue[i] = problem{
			question: record[0],                    // remember that each record is a 2 element slice [...] that represents a `question, answer` pair
			answer:   strings.TrimSpace(record[1]), // strings.TrimSpace() helps to remove spaces around answers from the CSV file
		}
	}
	return returnedValue
}

// openFile() uses the system to open and prep the file for reading
func openFile() *os.File {
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

func main() {
	// parse flags
	flag.Parse() // required to initialize the specified flags with the Operating system
}
