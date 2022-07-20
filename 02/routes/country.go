package routes

import (
	"net/http"
	"packages/controllers"
)

func CountryRoutes() {
	http.HandleFunc("/country", controllers.CountryIndex)
}
