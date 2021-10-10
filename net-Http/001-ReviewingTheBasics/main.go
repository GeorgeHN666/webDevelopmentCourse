package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "whatever")
}
func main() {
	var j hotdog
	http.ListenAndServe(":8080", j)
}
