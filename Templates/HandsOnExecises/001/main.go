package main

import (
	"log"
	"os"
	"text/template"
)

type Course struct {
	Number, Name, Units string
}
type semester struct {
	Term    string
	Courses []Course
}
type year struct {
	AcaYear              string
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}
func main() {
	years := year{
		AcaYear: "2019-2020",
		Fall: semester{
			Term: "Fall",
			Courses: []Course{
				Course{"chr3", "Golang For NLP", "4"},
				Course{"chr3", "Golang For Tokenization", "5"},
				Course{"chr3", "Golang For ML", "7"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, years)
	if err != nil {
		log.Fatal(err)
	}

}
