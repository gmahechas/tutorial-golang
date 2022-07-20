package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	msg := "Hello" + r.Host
	fmt.Fprintf(w, msg)
}

func main() {

	http.HandleFunc("/", index)

	fmt.Println("App is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
