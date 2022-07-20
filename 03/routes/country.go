package routes

import (
	"packages/controllers"

	"github.com/gorilla/mux"
)

func CountryMuxRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.CountryIndex).Methods("GET")
	return router
}
