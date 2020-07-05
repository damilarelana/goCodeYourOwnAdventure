package middleware

// // multiFlagTester()
// //  * checks to see if the user is trying to use multiple file formats (yaml, json, sql) at the same time
// //  * check to set number of flags `n` that have been set
// //	* throws an error message, if `n` > 1
// //  * returns a true boolean, if `n` > 1
// func multiFlagTester() bool {
// 	numFlag := flag.NFlag()
// 	if numFlag > 1 {
// 		return true
// 	}
// 	return false
// }

// // formatHandler()
// // * validates to avoid multiple flags being used simultaneously
// // * checks for which flag is being used
// // * leverages the appropriate flag handler
// // * defaults to using yaml flag when no flag is chosen
// func selectFlagHandler(mapHandler http.HandlerFunc) http.HandlerFunc {
// 	if multiFlagTester() { // check if multiple flags are being used i.e. multiFlagTester returns a true or false boolean
// 		errMsgHandler(fmt.Sprintf("Cannot use multiple flags at once. Please choose only yaml or json or sql \n"), nil)
// 	}

// 	if flag.NFlag() != 0 && !reflect.DeepEqual(*jsonFilename, "") { // if json filename is used THEN yaml/sql flag would be empty string
// 		jsonHandler := jsonFlagHandler(jsonFilename, mapHandler)
// 		fmt.Printf("Now using the JSON flag with the file: %s\n", *jsonFilename)
// 		return jsonHandler
// 	}

// 	if flag.NFlag() != 0 && !reflect.DeepEqual(*yamlFilename, "") { // if yaml filename is used THEN json/sql flags would be empty string
// 		yamlHandler := yamlFlagHandler(yamlFilename, mapHandler)
// 		return yamlHandler
// 	}

// 	if flag.NFlag() != 0 && !reflect.DeepEqual(*sqlDatabasePath, "") { // if sql flag is used THEN yaml/json flags would be empty string
// 		sqlHandler := sqlFlagHandler(sqlDatabasePath, mapHandler)
// 		return sqlHandler
// 	}

// 	// defaults to yamlHandler when no flags are selected i.e. 'flag.NFlag() returns 0'
// 	*yamlFilename = "pathsData.yaml"
// 	fmt.Printf("No file flags was set. Defaulting to file: %s\n", *yamlFilename)
// 	yamlHandler := yamlFlagHandler(yamlFilename, mapHandler)
// 	return yamlHandler
// }
