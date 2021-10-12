package main

import "net/http"

func main() {
	http.ListenAndServe(":8060", http.FileServer(http.Dir(".")))
}

// You Have To Have A File Call Index.html To Make It Run
// If You Have The File It Will Show The Page
// If You Dont Have Name It Correctly You Will Have The Root, Showing All The Documents In That Directory
