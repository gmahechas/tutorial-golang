package main

import (
	"fmt"
	"net/http"
	"packages/routes"
)

func main() {
	routes.CountryRoutes()
	routes.EstateRoutes()
	fmt.Print("App is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
