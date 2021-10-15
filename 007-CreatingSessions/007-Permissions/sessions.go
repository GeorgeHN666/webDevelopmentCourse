package main

import "net/http"

func getUser(r *http.Request) user {
	var u user
	c, err := r.Cookie("session")
	if err != nil {
		return u
	}
	// If Alredy Exist
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alredyLoggedIn(r *http.Request) bool {
	// Get Cookie
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
