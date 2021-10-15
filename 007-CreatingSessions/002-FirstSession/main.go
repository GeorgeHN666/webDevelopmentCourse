package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template

var dbUsers = map[string]user{} //User ID, user

var dbSessions = map[string]string{} //Session ID , session

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8055", nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	// Get The Cookie
	c, err := r.Cookie("sessions")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "sessions",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// If The User Exist Alredy ,Get The User Id
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[c.Value] = un
		dbUsers[un] = u
	}

	tpl.ExecuteTemplate(w, "index.html", u)
}
func bar(w http.ResponseWriter, r *http.Request) {
	// get Cookie
	c, err := r.Cookie("sessions")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.html", u)
}
