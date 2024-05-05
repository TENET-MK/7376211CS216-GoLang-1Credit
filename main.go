package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "TENET")
	})

	http.HandleFunc("/Register_Number", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376211CS216")
	})

	http.HandleFunc("/Department", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Computer science and Engineering")
	})

	http.HandleFunc("/Color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Black")
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/hello\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
