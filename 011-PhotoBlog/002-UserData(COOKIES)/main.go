package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8888", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	tpl.ExecuteTemplate(w, "index.html", c.Value)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:     "session",
			Value:    sID.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, c)
	}
	return c
}
