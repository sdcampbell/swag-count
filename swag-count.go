package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	urlPtr := flag.String("u", "", "URL to the swagger.json")
	filePtr := flag.String("f", "", "Path to the swagger.json file")

	flag.Parse()

	if (*urlPtr == "" && *filePtr == "") || (*urlPtr != "" && *filePtr != "") {
		fmt.Println("Usage: go run main.go -u [URL to swagger.json] or -f [path to swagger.json]")
		os.Exit(1)
	}

	var fileContent []byte
	var err error

	if *urlPtr != "" {
		// Fetch the Swagger JSON from the provided URL
		resp, err := http.Get(*urlPtr)
		if err != nil {
			fmt.Println("Error fetching URL:", err)
			os.Exit(1)
		}
		defer resp.Body.Close()

		fileContent, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			os.Exit(1)
		}
	} else {
		// Read the file content
		fileContent, err = ioutil.ReadFile(*filePtr)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
	}

	// Decode the JSON content
	var swagger map[string]interface{}
	if err := json.Unmarshal(fileContent, &swagger); err != nil {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	// Count the number of endpoints
	endpointCount := 0
	if paths, ok := swagger["paths"].(map[string]interface{}); ok {
		for _, methods := range paths {
			for range methods.(map[string]interface{}) {
				endpointCount++
			}
		}
	}

	fmt.Printf("Number of API endpoints: %d\n", endpointCount)
}
