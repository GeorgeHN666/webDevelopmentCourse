package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template/index.html"))
}

// type person struct{
// 	FirstName string
// 	LastName  string
// 	Subscribe bool
// }

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	data := make([]byte, r.ContentLength)
	r.Body.Read(data)
	body := string(data)

	err := tpl.ExecuteTemplate(w, "index.html", body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
