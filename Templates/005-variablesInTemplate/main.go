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
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", "Self-Confidence, Self-Focus,Self-Learning")
	if err != nil {
		log.Fatal(err)
	}
}

// In This Example Were Gonna Look How To Assign A Value To A Template And PassinG iNTO A Template
