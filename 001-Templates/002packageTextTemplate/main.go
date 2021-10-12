package main

import (
	"log"
	"os"
	"text/template"
)

// The Important Thing Is To Know How To Parse The Documents And Also Execute them!!
func main() {
	tpl, err := template.ParseFiles("tpl.html")
	if err != nil {
		log.Fatal(err)
	}

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Could't Creat The File: LN.15", err)
	}
	defer nf.Close()

	err = tpl.Execute(nf, nil)
	if err != nil {
		log.Fatal(err)
	}
}
