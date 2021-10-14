package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template/index.html"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8066", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	var s string
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		// Open it
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		// Read it
		read, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		s = string(read)
		// Store it
		store, err := os.Create(filepath.Join("./users/george", h.Filename))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer store.Close()
		_, err = store.Write(read)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl.ExecuteTemplate(w, "index.html", s)
}
