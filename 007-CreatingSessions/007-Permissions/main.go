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
	Role     string
}

var tpl *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8060", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	if !alredyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if u.Role != "007" {
		http.Error(w, "You Must Be 007 To Enter The Bar", http.StatusForbidden)
		return
	}

	tpl.ExecuteTemplate(w, "bar.html", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	// If Alredy Sign up
	if alredyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var u user

	// Process Post
	if r.Method == http.MethodPost {

		// Get Data

		// Check If Name Alredy In Use

		// Create Session

		// Store Values In dbUser

		// Redirect

		// Get
		e := r.FormValue("email")
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		rol := r.FormValue("role")

		// Check
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Name Alredy Taken", http.StatusForbidden)
			return
		}

		// Create S
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		// Store
		cp, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		u = user{e, un, cp, f, l, rol}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "signup.html", u)
}

func login(w http.ResponseWriter, r *http.Request) {
	// Check If Alredy Logged In
	if alredyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	var u user

	// Process
	if r.Method == http.MethodPost {
		e := r.FormValue("email")
		p := r.FormValue("password")

		// Check If Theres A Email
		em, ok := dbUsers[e]
		if !ok {
			http.Error(w, "email or/and Password Don't Match", http.StatusForbidden)
			return
		}

		// See If The Stored Password Patch
		err := bcrypt.CompareHashAndPassword(em.Password, []byte(p))
		if err != nil {
			http.Error(w, "email or/and Password Don't Match", http.StatusForbidden)
			return
		}

		// Create Session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = e

		// Redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", u)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alredyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Get Cookie
	c, _ := r.Cookie("session")
	// Delete Session
	delete(dbSessions, c.Value)
	// Delete Cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
