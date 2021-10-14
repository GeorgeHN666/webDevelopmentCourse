package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8000", nil)
}
func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post">
	<input type="text" name="q">
	<input type="submit" placeholder="submit">
	</form>
	<br>`+v)
	// IF We Change The Method To Post It Will Send The Response Throught The Body
	// IF We Change The Method To Get It Will Send The Data Throught The URL But Also Throught The Body In This Case Cause We're Concatinating
}
