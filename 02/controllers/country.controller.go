package controllers

import (
	"fmt"
	"net/http"
)

func CountryIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world from country index")
}
