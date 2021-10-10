package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dooooooggggg")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Caaaaaaaaaatttt")
}

func main() {
	// http.Handle Does't Have The Handler Method Attached So In Order To Pass It As A Handler We Have To Use Conversion With HTTP.HandlerFunc Or Just Use HandleFunc That Have The Handler Method Attached
	// Don't Get Confused Bettwen ***HandleFunc*** and ***HandlerFunc***
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))
	// http.HandleFunc("/dog/",d)
	// http.HandleFunc("/dog/",d)

	http.ListenAndServe(":8000", nil)
}
