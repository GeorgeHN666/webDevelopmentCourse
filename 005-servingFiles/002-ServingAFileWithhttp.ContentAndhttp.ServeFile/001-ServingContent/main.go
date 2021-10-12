package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", cuzco)
	http.HandleFunc("/cuzco", cuzcopic)
	http.ListenAndServe(":8030", nil)
}
func cuzco(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="./resources/cuzco.jpg">`)
}
func cuzcopic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./resources/cuzco.jpg")
	if err != nil {
		http.Error(w, "File Not Found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "File Not Found", 404)
		return
	}
	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
}
