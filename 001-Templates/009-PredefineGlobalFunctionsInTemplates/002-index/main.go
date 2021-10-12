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
	xS := []string{"Zero", "One", "Two", "Three", "Four", "Five"}

	data := struct {
		Words []string
		Lname string
	}{
		xS,
		"Hernandez",
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
