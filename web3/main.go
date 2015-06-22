package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

type Page struct {
	Title string
	Body  string
}

func handler(w http.ResponseWriter, r *http.Request) {

	path, err := filepath.Abs(filepath.Join(basePath(), r.URL.Path))
	if err != nil {
		panic(err)
	}

	if file, err := ioutil.ReadFile(path); err != nil {
		log.Printf("Error loading %v: %v", path, err.Error())
		w.WriteHeader(http.StatusNotFound)
	} else {
		log.Printf("Serving you %v", path)
		if t, err := template.New(path).Parse(string(file)); err != nil {
			log.Printf("Error loading template: %v", err)
		} else {
			t.Execute(w, Page{path, "<Test>"})
		}
	}
}

func basePath() string {
	if cwd, err := filepath.Abs("."); err == nil {
		return cwd
	}
	panic("D'oh!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
