package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"layres_new/delivery"
	"layres_new/services"
	"layres_new/store"
	"log"
	"net/http"
)

func main() {
	var db, err = sql.Open("mysql", "root:Manish@123Sharma@/Customer_services")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)// proper error handling instead of panic in your app
	}
	r := mux.NewRouter()
	datastore := store.New(db)
	defer datastore.CloseDB()
	service := services.New(datastore)
	handler := delivery.New(service)

	r.HandleFunc("/customer",handler.GetByName).Methods(http.MethodGet).Queries("name","{name}")
	r.HandleFunc("/customer",handler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id}", handler.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/customer",handler.Create).Methods(http.MethodPost)
	r.HandleFunc("/customer/{id:[0-9]+}",handler.Remove).Methods(http.MethodDelete)
	r.HandleFunc("/customer/{id:[0-9]+}",handler.Update).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(":3005", r))
}
