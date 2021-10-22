package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/webDevelopmentCourse/013-Go_And_MongoDB/002-json/models"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	// added router plus teh parameter
	r.GET("/user/:id", getUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta http-equiv="X-UA-Compatible" content="IE=edge">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>BAR</title>
		<style>
			html{
				background-color: rgb(26, 25, 25);
				font-family: Georgia, 'Times New Roman', Times, serif;
				color: rgb(196, 201, 247);
				font-size: medium;
			}
		</style>
	</head>
	<body>

	</body>
	</html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "Male",
		Age:    35,
		Id:     p.ByName("id"),
	}

	// Marshal Into JSON
	uj, _ := json.Marshal(u)

	// Write Content Type, Statuscode, Payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) //200
	fmt.Fprintf(w, "%s\n", uj)
}
