package controllers

import (
	"fmt"
	"net/http"
)

func CountryIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from Country!")
}
