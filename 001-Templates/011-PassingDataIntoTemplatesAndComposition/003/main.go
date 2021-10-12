package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}
type semester struct {
	Term    string
	Courses []course
}
type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	y := year{Fall: semester{
		Term: "Fall",
		Courses: []course{
			{"CHR1", "Advance Go", "5"},
			{"CHR1", "Go 4 Data Science", "10"},
			{"CHR1", "Go For Begginers", "5"},
		},
	},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatal(err)
	}
}
