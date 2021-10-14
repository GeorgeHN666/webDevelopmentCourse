package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	io.WriteString(w, "Do My Search:"+v)
}

// To SEE The Results It This Execercise You Have To Put ?q=whateveryouwant
// ? stands For Query
// q Is the identifier that we chose

// ---------------------------------------------------------------------------------

// If You Use The POST Form It Will Pass The Values Throught The Body

// If You Use The Get Form It Will Pass The Values Throught The Url

/* Example Of The URL Values
http://video.google.co.uk:80/videoplay?docid=7246927612831078230&hl=en#00h02m30s
	http:// - It The Protocol
	video - Subdomain
	google.co.uk - Domain
	80 - Port
	/videoplay - Path
	? - Query
	docid=7246927612831078230&hl=en - Parameters
	#00h02m30s - Fragments
*/
