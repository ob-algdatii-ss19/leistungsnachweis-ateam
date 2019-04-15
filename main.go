package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

/**
struct for parsing JSON-objects
*/
type json_struct struct {
	Firstname string
	Lastname  string
}

/**
Interface for JSON-Requests
*/
func jsonInterfaceHandler(w http.ResponseWriter, req *http.Request) {

	//read json-data from request
	decoder := json.NewDecoder(req.Body)
	var receivedData json_struct
	err := decoder.Decode(&receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println("[INFO] Received JSON-Data: " + receivedData.Firstname + ", " + receivedData.Lastname)

	//TODO call go-method with received json (e.g. go-method with any graph-algorithms)

	//build json-data for response
	responseData := json_struct{"Max", "Huber"}
	responseDataString, err := json.Marshal(responseData)

	fmt.Print("[INFO] send JSON to frontend: ")
	fmt.Println(bytes.NewBuffer(responseDataString))

	if err != nil {
		panic(err)
	}
	_, _ = fmt.Fprintf(w, "%s", bytes.NewBuffer(responseDataString))

}

/**
Handler for delivering http-requests
*/
func viewHandler(w http.ResponseWriter, r *http.Request) {

	filePath := "frontend"

	//error handling for request to root-level
	if r.URL.Path == "/" {
		filePath += "/index.html"
	} else {
		filePath += r.URL.Path
	}

	fmt.Println(">>>> Requesting: " + r.URL.Path)

	file, err := os.Open(filePath)

	// error response with file does not exist
	if err != nil {
		fmt.Printf("%s not found\n", filePath)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<html><body style='font-size:2em'>four-oh-four. Could not find: "+filePath+"</body></html>")
		return
	}
	defer file.Close()
	fileStat, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
	}

	var fileType = strings.Split(filePath, ".")[1]
	_, filename := path.Split(filePath)
	t := fileStat.ModTime()

	fmt.Println("\t\t > Serving: " + filePath + " with filetype " + fileType)
	fmt.Printf("\t\t%v -- last modified-time:  %+v\n", filename, t)

	//change mime-type for css-files
	if fileType == "css" {
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	}

	http.ServeContent(w, r, filename, t, file)
}

func main() {
	//register handlers
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/json", jsonInterfaceHandler)

	//start server
	fmt.Println("Server started on port 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
