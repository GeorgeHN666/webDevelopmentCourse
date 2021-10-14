package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":3000", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">set a cookie</a></h1>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "some-Value",
	})
	fmt.Fprintln(w, `<h1><a href="/read">read</a></h1>`)
}
func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, c)
}

func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	c.MaxAge = -1 //delete cookie
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
