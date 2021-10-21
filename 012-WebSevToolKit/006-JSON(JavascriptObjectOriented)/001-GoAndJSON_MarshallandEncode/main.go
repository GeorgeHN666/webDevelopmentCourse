package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name  string
	Last  string
	Stuff []string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	http.Handle("/favincon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8050", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>FOO</title>
	</head>
	<body>
		you are at foo
	</body>
	</html>`
	w.Write([]byte(html))
}

func mshl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	p1 := person{"james", "bond", []string{"suit", "Gun", "Black Sense Of Humor"}}

	json, err := json.Marshal(p1)
	if err != nil {
		log.Panic(err)
	}
	w.Write(json)
}

func encd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "aplication/json")
	p1 := person{"james", "bond", []string{"suit", "Gun", "Black Sense Of Humor"}}

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Panic(err)
	}
}
