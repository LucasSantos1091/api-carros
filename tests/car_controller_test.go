package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"car-api/controllers"
	"car-api/models"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cars", controllers.GetCars).Methods("GET")
	r.HandleFunc("/cars/{id}", controllers.GetCarByID).Methods("GET")
	r.HandleFunc("/cars", controllers.CreateCar).Methods("POST")
	r.HandleFunc("/cars/{id}", controllers.UpdateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", controllers.DeleteCar).Methods("DELETE")
	return r
}

// Testes bem cobertos
func TestGetCars(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cars", nil)
	response := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}
}

func TestGetCarByID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cars/1", nil)
	response := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}
}

func TestCreateCar(t *testing.T) {
	car := []byte(`{"make":"Ford", "model":"Focus", "year":2022}`)
	req, _ := http.NewRequest("POST", "/cars", bytes.NewBuffer(car))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}
}

func TestUpdateCar(t *testing.T) {
	car := []byte(`{"make":"Ford", "model":"Focus", "year":2023}`)
	req, _ := http.NewRequest("PUT", "/cars/1", bytes.NewBuffer(car))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}
}

func TestDeleteCar(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/cars/1", nil)
	response := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", response.Code)
	}
}

// Testes mal cobertos
func TestInvalidPath(t *testing.T) {
	req, _ := http.NewRequest("GET", "/invalid", nil)
	response := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(response, req)
}

func TestEmptyRequest(t *testing.T) {
	req, _ := http.NewRequest("POST", "/cars", nil)
	response := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(response, req)
}

func TestInvalidMethod(t *testing.T) {
	req, _ := http.NewRequest("PATCH", "/cars", nil)
	response := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(response, req)
}

func TestInvalidCarID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cars/999", nil)
	response := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(response, req)
}

func TestEmptyCarsList(t *testing.T) {
	models.Cars = []models.Car{}
	req, _ := http.NewRequest("GET", "/cars", nil)
	response := httptest.NewRecorder()
	router := setupRouter()
	router.ServeHTTP(response, req)
}
