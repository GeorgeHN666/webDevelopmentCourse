package main

import "net/http"

func getUser(r *http.Request) user {
	var u user
	// get Cookies
	c, err := r.Cookie("session")
	if err != nil {
		return u
	}

	// if the use alredy exist,Get User
	if un, ok := dbSession[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alredyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSession[c.Value]
	_, ok := dbUsers[un]
	return ok
}
