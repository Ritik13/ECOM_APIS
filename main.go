package main

import (
	"ecom/controllers"
	"fmt"
	"net/http"

	"ecom/database"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Ecom Apis build on go, No framework used")
	database.Initialize()
	router := mux.NewRouter()

	// auth routes
	router.HandleFunc("/users/login", controllers.UserLogin).Methods("POST")
	router.HandleFunc("/users/register", controllers.RegisterUser).Methods("POST")

	// product routes
	router.HandleFunc("/products", controllers.AddNewProdect).Methods("POST")
	router.HandleFunc("/products", controllers.GetAllProducts).Methods("GET")
	router.HandleFunc("/products", controllers.UpdateProduct).Methods("PUT")
	http.ListenAndServe(":8080", router)
}
