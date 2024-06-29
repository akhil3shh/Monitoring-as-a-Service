package routes

import (
	"github.com/Pratham-Karmalkar/controller"
	"github.com/gorilla/mux"
)

var AppRoutes = func(router *mux.Router) {
	router.HandleFunc("/gen/{stat}", controller.StartGeneration).Methods("POST")
	//router.HandleFunc("/stop", controller.StopGeneration).Methods("GET")
}
