package main

import (
	"ecom/controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Ecom Apis build on go, No framework used")

	router := mux.NewRouter()

	// auth routes
	router.HandleFunc("/users/login", controllers.UserLogin).Methods("POST")

	http.ListenAndServe(":8080", router)
}
