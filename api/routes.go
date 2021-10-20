package api

import "github.com/gorilla/mux"

// routes wires Api handlers to mux routes
func routes(router *mux.Router, api *Api) {
	router.HandleFunc("/greeting", api.handleGreet())
	router.HandleFunc("/weather", api.handleForecast())
}
