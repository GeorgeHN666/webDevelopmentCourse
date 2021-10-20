package main

import (
	"html/template"
	"net/http"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8060", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c := Cookie(w, r)
	c = append(w, c)
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.html", xs)
}

func Cookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
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

func append(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	// Values
	p1 := "sunset.jpg"
	p2 := "monday.jpg"
	p3 := "happy.jpg"

	s := c.Value
	if !strings.Contains(s, p1) {
		s += "|" + p1
	}
	if !strings.Contains(s, p2) {
		s += "|" + p2
	}
	if !strings.Contains(s, p3) {
		s += "|" + p3
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
