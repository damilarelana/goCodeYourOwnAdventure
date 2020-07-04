package middleware

import (
	"encoding/json"
	"fmt"
	"os"

	cyoa "github.com/damilarelana/goCYOA"
)

// defines the error message handler
func errMsgHandler(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

// openFile ...
// uses the system to open and prep the file for reading
func openFile(storyFilename *string) *os.File {
	openedFile, err := os.Open(*storyFilename)
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to open JSON file %s : %s\n", *storyFilename, err.Error()))
	}
	return openedFile
}

// readFile ...
//  * takes the pointer to the opened file i.e. openedFile
func readFile(f *os.File) *json.Decoder {
	fileData := json.NewDecoder(f)
	return fileData
}

// parseFile parses the jsonData into structData
func parseFile(fileData *json.Decoder, story *cyoa.Story) {
	err := fileData.Decode(&story) // decode and parse the `json data` into `structs data` within the address space &story
	if err != nil {
		errMsgHandler(fmt.Sprintf("Failed to parse file data %v\n", *fileData))
	}
}

// JSONFileHandler ...
//	- pointer to the json filename : storyFilename
//	- pointer to the story struct: story
func JSONFileHandler(storyFilename *string, story *cyoa.Story) {
	openedFile := openFile(storyFilename) // open file
	fileData := readFile(openedFile)      // read file content
	parseFile(fileData, story)            // parse file data and store it in the memory address `story`
}
