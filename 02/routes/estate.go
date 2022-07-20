package routes

import (
	"fmt"
	"net/http"
)

func EstateRoutes() {
	http.HandleFunc("/estate", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world from estate")
	})
}
