package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/webDevelopmentCourse/013-Go_And_MongoDB/003-PostDelete/models"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	// added router
	r.POST("/user", createUser)
	// added router plus parameter
	r.DELETE("/user/:id", deleteUser)
	http.ListenAndServe("localhost:8050", r)

}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Index</title>
	</head>
	<body>
	<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
	</body>
	</html>`

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "Jorge Enrique",
		Gender: "Male",
		Age:    19,
		Id:     p.ByName("id"),
	}

	// Marshal Into JSON
	j, err := json.Marshal(u)
	if err != nil {
		log.Panic(err)
	}

	// Write Content-Type, StatusCode, Payload
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", j)
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Composite Literal - type and curly braces
	u := models.User{}

	// encode/decode for sending/receiving JSON to/from a stream
	json.NewDecoder(r.Body).Decode(&u)

	// change id
	u.Id = "007"

	// marshal/unmarshal for having JSON assigned to a variable
	j, _ := json.Marshal(u)

	// write content-Type , statuscode, payload
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", j)
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: write code to delete user
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Write Code to delete user\n")
}
