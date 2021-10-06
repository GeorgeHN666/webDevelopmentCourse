package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Persons struct {
	Index int
	Name  string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	p1 := Persons{
		Index: 1,
		Name:  "Andres Mejia",
	}
	p2 := Persons{
		Index: 2,
		Name:  "Beatriz Meza",
	}
	p3 := Persons{
		Index: 3,
		Name:  "Bruco",
	}
	p4 := Persons{
		Index: 4,
		Name:  "STN",
	}

	favoritePersons := []Persons{p1, p2, p3, p4}

	err := tpl.Execute(os.Stdout, favoritePersons)
	if err != nil {
		log.Fatal(err)
	}
}
