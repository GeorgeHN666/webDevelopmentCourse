package main

import (
	"log"
	"os"
	"text/template"
)

type names struct {
	Name  string
	Value string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	p1 := names{
		Name:  "STN",
		Value: "In Nomine Dei Nostri",
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatal(err)
	}
}
