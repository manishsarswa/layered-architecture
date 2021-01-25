package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"layres/delivery"
	"layres/services"
	"layres/store"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	datastore := store.New()
	defer datastore.CloseDb()
	service := services.New(datastore)
	handler := delivery.New(service)

	r.HandleFunc("/customer",handler.GetCustomerByName).Methods(http.MethodGet).Queries("name","{name}")
	r.HandleFunc("/customer",handler.GetCustomer).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id}", handler.GetCustomerById).Methods(http.MethodGet)
	r.HandleFunc("/customer",handler.CreateCustomer).Methods(http.MethodPost)
	r.HandleFunc("/customer/{id:[0-9]+}",handler.RemoveCustomer).Methods(http.MethodDelete)
	r.HandleFunc("/customer/{id:[0-9]+}",handler.UpdateCustomer).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":3004", r))
}
