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

type user struct {
	Name  string
	Age   int
	Admin bool
}

func main() {
	u1 := user{Name: "George", Age: 19, Admin: true}
	u2 := user{Name: "Cuzco", Age: 105, Admin: true}
	u3 := user{Name: "Beatriz", Age: 47, Admin: false}
	u4 := user{Name: "Borrego", Age: 45, Admin: false}

	users := []user{u1, u2, u3, u4}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatal(err)
	}
}
