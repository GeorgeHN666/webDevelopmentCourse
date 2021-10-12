package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("George-Key", "This Is From George")
	r.Header.Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1> This Is A Sample Text</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8000", d)
}
