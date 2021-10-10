package main

import (
	"io"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "The Dogs Are So Adorable")
	case "/cat":
		io.WriteString(w, "Cats Are So Curious And Adorable")
	default:
		io.WriteString(w, "This Is The Default Option")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}

// This Is Just A Example Of How A Multiplex Server Works
