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
	xSnames := []string{"Gandhi", "MLK", "Buddha", "Jesus", "Muhammad", "STN"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", xSnames)
	if err != nil {
		log.Fatal(err)
	}
}
