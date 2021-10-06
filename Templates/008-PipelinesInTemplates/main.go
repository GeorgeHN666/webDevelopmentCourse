package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func double(f float64) float64 {
	return f + 2
}
func mult(f float64) float64 {
	return f * 2
}
func div(f float64) float64 {
	return f / 2
}
func power(f float64) float64 {
	return f * f
}
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

var fm = template.FuncMap{
	"Db": double,
	"Mt": mult,
	"D":  div,
	"P":  power,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 5.0)
	if err != nil {
		log.Fatal(err)
	}
}
