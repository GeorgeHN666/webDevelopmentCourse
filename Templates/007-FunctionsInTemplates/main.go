package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	// tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	tpl = template.Must(template.New("tpl.gohtml").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type Person struct {
	Name string
}

func main() {
	p1 := Person{
		Name: "John",
	}
	p2 := Person{
		Name: "Michael",
	}
	p3 := Person{
		Name: "James",
	}
	p4 := Person{
		Name: "Rochell",
	}
	Peoples := []Person{p1, p2, p3, p4}

	err := tpl.Execute(os.Stdout, Peoples)
	if err != nil {
		log.Fatal(err)
	}
}
