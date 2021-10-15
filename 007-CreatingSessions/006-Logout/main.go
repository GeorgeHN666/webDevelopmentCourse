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
	// Were Creating A User
	pass, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["test666"] = user{"test@gmail.com", "test666", pass, "john", "macleod"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8888", nil)
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
	tpl.ExecuteTemplate(w, "bar.html", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	// Check If Alredy Sign Up
	if alredyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	var u user

	// Process Submition
	if r.Method == http.MethodPost {
		// Get Data Form Value
		e := r.FormValue("email")
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		// Check If Name Alredy Taken
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
		// Store Values In dbUser
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

func login(w http.ResponseWriter, r *http.Request) {
	// Check If Alredy Signup
	if alredyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u user
	// Proccess Form Submition
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")

		// is there a Username??
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or Password Don't Match", http.StatusForbidden)
			return
		}

		// Does The Password Store Match??
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or Password Don't Match", http.StatusForbidden)
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "login.html", u)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alredyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	c, _ := r.Cookie("session")
	// delete Session
	delete(dbSession, c.Value)
	// remove cookie
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
