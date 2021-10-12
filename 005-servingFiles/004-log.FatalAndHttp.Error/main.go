package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}

// Log.Fatal Will Catch The Error And Will Print It In StandarOut (Stdout) And Calls os.Exit()

// http.Error Will Trow The Error As A Plain Text That We Can Modified And Will Give The Error Code ej:(404)
