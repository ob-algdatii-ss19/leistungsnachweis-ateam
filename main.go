package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Body []byte
}

type test_struct struct {
	Firstname string
	Lastname  string
}

func loadPage(title string) (*Page, error) {
	filename := "frontend/" + title

	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	println("Requesting: " + r.URL.Path)

	var p *Page
	var err error

	if r.URL.Path == "/" {
		p, err = loadPage("index.html")
	} else {
		p, err = loadPage(r.URL.Path)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) //page not found
		return
	}
	_, _ = fmt.Fprintf(w, "%s", p.Body)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	p, err := loadPage("result")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) //page not found
		return
	}
	_, _ = fmt.Fprintf(w, "%s", p.Body)
}

func validationHandler(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	_, _ = fmt.Fprintf(w, "%s", "{\"response\": \"firstname:"+t.Firstname+"; lastname:"+t.Lastname+"\"}")
}

func main() {
	fmt.Println("Server started on port 8080 ...")

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/validate", validationHandler)
	http.HandleFunc("/result", resultHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
