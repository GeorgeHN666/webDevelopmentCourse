package main

import (
	"html/template"
	"log"
	"os"
)

type page struct {
	Title, Heading, Input string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

// HTML Template Scape All The Dangerous Characters Are Skipped
func main() {
	home := page{"Escaped", "Danger Is Scaping With HTML/TEMPLATES", `<script>alert("Yow!");</script>`}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", home)
	if err != nil {
		log.Fatal(err)
	}
}
