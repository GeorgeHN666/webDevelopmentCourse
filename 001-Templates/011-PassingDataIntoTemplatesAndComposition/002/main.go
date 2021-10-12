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
type doubleZero struct {
	person
	Ltk bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}
func main() {
	p1 := doubleZero{person{"Agent", 47}, false}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatal(err)
	}
}
