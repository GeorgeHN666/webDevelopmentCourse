package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Email    string
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSession = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8066", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	if !alredyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.html", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	// Check If Its Alredy Logged In
	if alredyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var u user

	// Process Submission
	if r.Method == http.MethodPost {

		// Get Values
		e := r.FormValue("email")
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		// Name Alredy Taken?
		if _, ok := dbSession[un]; ok {
			http.Error(w, "Name Alredy Taken", http.StatusForbidden)
			return
		}
		// Create Session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSession[c.Value] = un
		// Store Values IN dbUsers
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		u = user{e, un, bs, f, l}
		dbUsers[un] = u
		// Redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.html", u)
}
