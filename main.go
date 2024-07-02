package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Building Rest APIs")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	mux.HandleFunc("GET /comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "return a single comment for comment with id: %s", id)
	})

	mux.HandleFunc("GET /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "return all comments")
	})

	mux.HandleFunc("POST /comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "post a new comment")
	})

	if err := http.ListenAndServe("localhost:8000", mux); err != nil {
		fmt.Println(err.Error())
	}
}
