package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.html", "Be FOCUS ")
	if err != nil {
		log.Fatal(err)
	}
}
