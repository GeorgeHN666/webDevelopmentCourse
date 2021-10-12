package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", bruco)
	http.HandleFunc("/bruco", brucopic)
	http.ListenAndServe(":8080", nil)
}

func bruco(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<img src="cusco.jpg"
	`)
}

func brucopic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("cusco.jpg")
	if err != nil {
		http.Error(w, "file Not Found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}
