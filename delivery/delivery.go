package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"layres/entities"
	"layres/services"
	"net/http"
	"strconv"
)

type Handler struct{
	service services.CustomerService
}


func New(customer services.CustomerService) Handler {
	return Handler{service: customer}
}
func (c Handler) GetCustomerById(w http.ResponseWriter,r *http.Request){

	vars:=mux.Vars(r)
	id, _ :=strconv.Atoi(vars["id"])
	c.service.GetCustomerById(w,id)
}

func (c Handler) GetCustomerByName(w http.ResponseWriter,r *http.Request){
	Name :=r.URL.Query().Get("name")
	c.service.GetCustomerByName(w,Name)
}


func (c Handler) CreateCustomer(w http.ResponseWriter,r *http.Request){
	var customer entities.Customer
	body,_:=ioutil.ReadAll(r.Body)
	json.Unmarshal(body,&customer)

	c.service.CreateCustomer(w,customer)

}
func (c Handler) GetCustomer(w http.ResponseWriter,r *http.Request){
	c.service.GetCustomer(w)
}

func (c Handler) RemoveCustomer(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	id, _ :=strconv.Atoi(vars["id"])
	c.service.RemoveCustomer(w,id)
}


func (c Handler) UpdateCustomer(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	id, _ :=strconv.Atoi(vars["id"])
	var customer entities.Customer
	body,_:=ioutil.ReadAll(r.Body)
	json.Unmarshal(body,&customer)
	resp,ok:=c.service.UpdateCustomer(customer,id)
	if ok!=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(resp)
}