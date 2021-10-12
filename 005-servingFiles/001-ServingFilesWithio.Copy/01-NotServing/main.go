package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8030", nil)
}

// Here We Are Serving From Another Source
func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;  charset=utf-8")

	io.WriteString(w, `
	<!--not Serving From Our Server-->
	<img src="https://i.natgeofe.com/n/8a3e578f-346b-479f-971d-29dd99a6b699/nationalgeographic_2751013_4x3.jpg">
	`)
}
