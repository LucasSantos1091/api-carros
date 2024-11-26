package main

import (
	"log"
	"net/http"

	"car-api/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cars", controllers.GetCars).Methods("GET")
	r.HandleFunc("/cars/{id}", controllers.GetCarByID).Methods("GET")
	r.HandleFunc("/cars", controllers.CreateCar).Methods("POST")
	r.HandleFunc("/cars/{id}", controllers.UpdateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", controllers.DeleteCar).Methods("DELETE")

	log.Println("API is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
