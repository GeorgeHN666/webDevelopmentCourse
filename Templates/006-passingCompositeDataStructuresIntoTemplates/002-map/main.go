package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	xSnames := map[string]string{
		"india":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Hate":     "Jesus",
		"Love":     "STN",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xSnames)
	if err != nil {
		log.Fatal(err)
	}
}
