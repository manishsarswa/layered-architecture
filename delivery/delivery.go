package delivery

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"layres_new/entities"
	"layres_new/errors"
	"layres_new/services"
	"net/http"
	"strconv"
)

type Handler struct{
	service services.CustomerService
}



func New(customer services.CustomerService) Handler {
	return Handler{service: customer}
}

//get customer by id will return all the detail of customer on the basis of id in the response
func (c Handler) GetByID(w http.ResponseWriter,r *http.Request){

	vars:=mux.Vars(r)
	id, _ :=strconv.Atoi(vars["id"])
	if id<=0 {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, string(errors.NotFound))
		return
	}

	resp,err:=c.service.GetByID(id)
	if err != nil{
		io.WriteString(w,err.Error())
		return
	}

	if resp.Id==0{
		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, string(errors.NotFound))
		return
	}

	json.NewEncoder(w).Encode(resp)


}

func (c Handler) GetByName(w http.ResponseWriter,r *http.Request){
	Name :=r.URL.Query().Get("name")
	resp,err:=c.service.GetByName(Name)
	if err != nil {
		io.WriteString(w,err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(resp)

}

//create customer will add new customers in the database
func (c Handler) Create(w http.ResponseWriter,r *http.Request){
	var customer entities.Customer
	body,_:=ioutil.ReadAll(r.Body)
	json.Unmarshal(body,&customer)

	resp,err:=c.service.Create(customer)
	if err != nil || resp.Id==0{
		io.WriteString(w,err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
//get customer will return details of all the customers
func (c Handler) GetAll(w http.ResponseWriter,r *http.Request){
	resp,err:=c.service.GetAll()
	if err != nil {
		io.WriteString(w,err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(resp)

}
//Remove customer will remove customer from database on the basis of id in the request
func (c Handler) Remove(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	id, _ :=strconv.Atoi(vars["id"])
	err:=c.service.Remove(id)

	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//update customer will update the details of customer other than id and dateOfBirth(dob)
func (c Handler) Update(w http.ResponseWriter,r *http.Request){
	vars:=mux.Vars(r)
	id, _ :=strconv.Atoi(vars["id"])
	var customer entities.Customer
	body,_:=ioutil.ReadAll(r.Body)
	json.Unmarshal(body,&customer)

	resp,err:=c.service.Update(customer,id)

	if err!=nil || resp.Id==0{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(resp)
}