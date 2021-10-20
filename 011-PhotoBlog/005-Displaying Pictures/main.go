package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8055", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	// Get The Cookie
	c := getCookie(w, r)

	// Conditional
	if r.Method == http.MethodPost {
		mf, fh, err := r.FormFile("upload_files")

		handleError(err)

		// Creayte Sha For File Name
		sP := strings.Split(fh.Filename, ".")[1]

		hash := sha1.New()

		io.Copy(hash, mf)
		fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + sP

		// Create File
		wd, err := os.Getwd()

		handleError(err)

		path := filepath.Join(wd, "public", "pics", fname)

		nf, err := os.Create(path)

		handleError(err)

		defer nf.Close()

		// Copy

		mf.Seek(0, 0)
		io.Copy(nf, mf)

		// Add File To The User Cookie
		c = appendValue(w, c, fname)
	}
	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.html", xs[1:])
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
