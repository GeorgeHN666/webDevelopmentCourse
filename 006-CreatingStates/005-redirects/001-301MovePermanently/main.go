package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your Request Method At Foo:", r.Method)
}
func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your Request Method At Bar")
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// This Status Code Will Store That Location At Your Browser And When Your Request That Path It Will Automaticly Redirect At That Location
