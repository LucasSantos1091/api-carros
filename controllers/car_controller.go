package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"car-api/models"

	"github.com/gorilla/mux"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.Cars)
}

func GetCarByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, car := range models.Cars {
		if car.ID == id {
			json.NewEncoder(w).Encode(car)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	json.NewDecoder(r.Body).Decode(&car)
	car.ID = len(models.Cars) + 1
	models.Cars = append(models.Cars, car)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, car := range models.Cars {
		if car.ID == id {
			json.NewDecoder(r.Body).Decode(&car)
			car.ID = id
			models.Cars[i] = car
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(car)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for i, car := range models.Cars {
		if car.ID == id {
			models.Cars = append(models.Cars[:i], models.Cars[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(models.Cars)
			return
		}
	}
	http.Error(w, "Car not found", http.StatusNotFound)
}
