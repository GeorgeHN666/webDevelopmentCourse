package main

import "net/http"

func getUser(r *http.Request) user {
	var u user
	// Get The Cookie
	c, err := r.Cookie("session")
	if err != nil {
		return u
	}

	// If Alredy Exist Get The Cookie
	if un, ok := dbSession[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alredyLoggedIn(r *http.Request) bool {
	// Get The Cookie
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSession[c.Value]
	_, ok := dbUsers[un]
	return ok
}
