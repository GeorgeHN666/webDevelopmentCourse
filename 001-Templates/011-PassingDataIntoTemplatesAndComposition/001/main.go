package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}
func main() {
	p1 := person{"John", 34}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatal(err)
	}
}
