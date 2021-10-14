package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8060", nil)
}
func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookieee",
		Value: "still some-value",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN")
	fmt.Fprintln(w, "Check Your Browser Cookies")
}
func read(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("my-cookie")
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1:", c1)
	}

	c2, err := r.Cookie("general")
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2", c2)
	}

	c3, err := r.Cookie("specific")
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #3", c3)
	}
}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "Some other value about general things",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "Some Other value about specific things",
	})

	fmt.Fprintln(w, "COOKIES WRITTEN")
	fmt.Fprintln(w, "In Your Browser Check The Cookies Status")
}
