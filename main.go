package main

import (
	"./backend"
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
Interface for JSON-Requests.

Detailed description of the interface:
{
	"Settings": {
		"Algorithm": enum //BASIC_GREEDY=0, WELSH_POWELL=1, BRON_KERBOSCH=2
	},
	"Intersection": {
		"Top": {
			"RightLane": boolean,
			"StraightLane": boolean,
			"LeftLane": boolean,
			"Pedestrian": enum //OFF=0, NORMAL= 1, WITH_ISLAND=2
		},
		"Right": {
			"RightLane": boolean,
			"StraightLane": boolean,
			"LeftLane": boolean,
			"Pedestrian": enum //OFF=0, NORMAL= 1, WITH_ISLAND=
		},
		"Buttom": {
			"RightLane": boolean,
			"StraightLane": boolean,
			"LeftLane": boolean,
			"Pedestrian": enum //OFF=0, NORMAL= 1, WITH_ISLAND=
		},
		"Left": {
			"RightLane": boolean,
			"StraightLane": boolean,
			"LeftLane": boolean,
			"Pedestrian": enum //OFF=0, NORMAL= 1, WITH_ISLAND=
		}
	}
}
*/
func jsonInterfaceHandler(w http.ResponseWriter, req *http.Request) {

	//read json-data from request
	decoder := json.NewDecoder(req.Body)
	var receivedData backend.GuiRequestData
	err := decoder.Decode(&receivedData)
	if err != nil {
		panic(err)
	}

	fmt.Println("[INFO] Received JSON-Data: ", receivedData)

	responseData := backend.HandleAlgorithmCalls(receivedData)

	//build json-data for response
	responseDataString, err := json.Marshal(responseData)

	fmt.Println("[INFO] Send JSON to frontend: ", responseData)

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

	//retrieve port for heroku
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	//start server
	fmt.Println("[INFO] Server started on port " + port + " ...")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
