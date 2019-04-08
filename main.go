package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
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

	var filetype = strings.Split(r.URL.Path, ".")[1]

	println("Requesting: " + r.URL.Path + " with filetype " + filetype)

	if filetype == "css" {
		cssHandler(w, r)
		return
	}

	var p *Page
	var err error

	if r.URL.Path == "/" {
		http.Redirect(w, r, "index.html", 301)
		return
		//p, err = loadPage("index.html")
	} else {
		p, err = loadPage(r.URL.Path)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) //page not found
		return
	}
	//http.ServeContent(w, r, p.Body, fi.ModTime(), f)

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

func cssHandler(w http.ResponseWriter, r *http.Request) {

	println("started cssHandler")

	filePath := "frontend" + r.URL.Path

	println("Requesting: " + filePath)

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("%s not found\n", filePath)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<html><body style='font-size:100px'>four-oh-four</body></html>")
		return
	}
	defer file.Close()
	fileStat, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Printf("serve %s\n", filePath)

	println("serve: " + filePath)

	_, filename := path.Split(filePath)
	t := fileStat.ModTime()
	fmt.Printf("time %+v\n", t)
	w.Header().Set("Content-Type", "text/css; charset=utf-8")
	http.ServeContent(w, r, filename, t, file)
}

func main() {
	fmt.Println("Server started on port 8080 ...")

	//http.HandleFunc("/", cssHandler)
	//http.Handle("/static/css", http.FileServer(http.Dir("frontend/static")))
	//http.HandleFunc("/validate", validationHandler)
	//http.HandleFunc("/result", resultHandler)
	http.HandleFunc("/", viewHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
