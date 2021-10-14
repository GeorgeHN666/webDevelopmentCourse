package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("template/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your Request Method At Foo:", r.Method)
}
func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your Requested Method At Bar:", r.Method)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your Requested Method At Barred:", r.Method)
	tpl.ExecuteTemplate(w, "index.html", nil)
}

// This Status Code Will Keep The Same Method That You Was Using
