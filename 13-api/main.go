package main

import (
	"fmt"
	"log"
	"net/http"
)

func hola(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola!")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(hola))
	/* mux.HandleFunc("/", hola) */
	mux.HandleFunc("/not-found", notFound)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Starting the server...")
	log.Fatal(server.ListenAndServe())
}
