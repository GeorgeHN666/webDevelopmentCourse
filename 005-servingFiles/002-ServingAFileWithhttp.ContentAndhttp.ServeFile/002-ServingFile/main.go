package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", bruco)
	http.HandleFunc("/bruco", brucopic)
	http.ListenAndServe(":8080", nil)
}

func bruco(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="cusco.jpg"`)
}
func brucopic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "cusco.jpg")
}
