package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Pratham-Karmalkar/routes"
	"github.com/gorilla/mux"
)

// MyData is a sample struct representing your JSON data

func main() {

	r := mux.NewRouter()
	routes.AppRoutes(r)
	http.Handle("/", r)
	fmt.Print("server running ...")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
