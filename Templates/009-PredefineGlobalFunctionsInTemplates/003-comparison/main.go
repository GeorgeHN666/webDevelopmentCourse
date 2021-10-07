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
	g1 := struct {
		Score1 int
		Score2 int
	}{
		72,
		10,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatal(err)
	}
}
