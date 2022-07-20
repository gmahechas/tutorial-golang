package main

import (
	"fmt"
	"net/http"
	"packages/routes"
)

func main() {
	router := routes.CountryMuxRoutes()
	fmt.Println("App started")
	http.ListenAndServe(":8080", router)
}
