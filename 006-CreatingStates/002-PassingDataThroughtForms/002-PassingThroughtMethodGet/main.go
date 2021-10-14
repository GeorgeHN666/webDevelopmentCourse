package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8055", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("g")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="get">
	<input type="text" name="g">
	<input type="submit" placeholder="submit">
	</form>
	<br>`+v)
}
