package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}
func main() {
	xS := []string{"zero", "One", "Two", "Three", "Four", "five"}

	err := tpl.Execute(os.Stdout, xS)
	if err != nil {
		log.Fatal(err)
	}
}
