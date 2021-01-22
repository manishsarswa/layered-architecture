package services

import (
	"encoding/json"
	"layres/entities"
	"layres/store"
	"net/http"
)

type CustomerService struct {
	store store.Customer
}

func New(customer store.Customer) CustomerService {
	return CustomerService{store: customer}
}

func (c CustomerService) GetCustomerById(w http.ResponseWriter, id int) {
	if id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := c.store.GetCustomerBYId(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]entities.Customer(nil))
	} else {
		if resp.Id == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode([]entities.Customer(nil))
		} else {
			json.NewEncoder(w).Encode(resp)
		}
	}
}

func (c CustomerService) GetCustomerByName(w http.ResponseWriter,Name string){
	resp,err:=c.store.GetCustomerByName(Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]entities.Customer(nil))
	} else {
			json.NewEncoder(w).Encode(resp)
	}

}

func (c CustomerService) CreateCustomer(w http.ResponseWriter,customer entities.Customer){

	age:=dateInSeconds(customer.Dob)
	if age/(365*24*3600)<18{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp,err:=c.store.CreateCustomer(customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]entities.Customer(nil))
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resp)
	}


}

func (c CustomerService) GetCustomer(w http.ResponseWriter){
	resp,err:=c.store.GetCustomer()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode([]entities.Customer(nil))
	} else {
		json.NewEncoder(w).Encode(resp)
	}
}


func (c CustomerService) RemoveCustomer(w http.ResponseWriter, id int){
	err:=c.store.RemoveCustomer(id)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c CustomerService) UpdateCustomer(customer entities.Customer,id int) (entities.Customer,error){
	cnt,err:=c.store.UpdateCustomer(customer,id)
	if err!=nil{
		return entities.Customer{},err
	}
	return cnt,nil
}
