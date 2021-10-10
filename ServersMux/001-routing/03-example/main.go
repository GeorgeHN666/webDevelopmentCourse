package main

import (
	"io"
	"net/http"
)

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "From Dog Function")
}

func c(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "From Cat Function")
}

func main() {
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}

// Here We Use The Default Server Mux
// This Code Structure Is The Best Practice For Mux Servers
