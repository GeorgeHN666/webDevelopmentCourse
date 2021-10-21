package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8050", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, "userID", 777)
	ctx = context.WithValue(ctx, "fname", "bond")

	results := dbAccess(ctx)

	fmt.Fprintln(w, results)
}

func dbAccess(ctx context.Context) int {
	uid := ctx.Value("userID").(int)

	return uid
}

func bar(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println(ctx)

	fmt.Fprintln(w, ctx)
}
