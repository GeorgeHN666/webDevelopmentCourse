package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}
type people struct {
	Name string
	Age  int
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	c1 := car{
		Manufacturer: "Ford",
		Model:        "Raptor",
		Doors:        4,
	}
	c2 := car{
		Manufacturer: "Ferrari",
		Model:        "F14",
		Doors:        2,
	}
	c3 := car{
		Manufacturer: "Pagganni",
		Model:        "Huayra",
		Doors:        2,
	}
	c4 := car{
		Manufacturer: "Tesla",
		Model:        "Model X",
		Doors:        4,
	}
	p1 := people{
		Name: "john",
		Age:  37,
	}
	p2 := people{
		Name: "Michael",
		Age:  42,
	}
	p3 := people{
		Name: "Karina",
		Age:  19,
	}
	p4 := people{
		Name: "Peter",
		Age:  45,
	}
	cars := []car{c1, c2, c3, c4}
	persons := []people{p1, p2, p3, p4}

	data := struct {
		Person []people
		Car    []car
	}{
		Person: persons,
		Car:    cars,
	}

	/* data := info{
		Person: persons,
		Car:    cars,
	}*/

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
