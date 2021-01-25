package services

import (
	"layres_new/entities"
	"layres_new/store"
)

type CustomerService struct {
	store store.Customer
}

func New(customer store.Customer) CustomerService {
	return CustomerService{store: customer}
}

func (c CustomerService) GetByID(id int) (entities.Customer,error){
	resp, err := c.store.GetByID(id)
	return resp,err
}

func (c CustomerService) GetByName(Name string) ([]entities.Customer,error){
	resp,err:=c.store.GetByName(Name)
	return resp,err

}

func (c CustomerService) Create(customer entities.Customer) (entities.Customer,error) {

	age := dateInSeconds(customer.Dob)

	if age/(365*24*3600) < 18 {
		return entities.Customer{},nil
	}
	return c.store.Create(customer)

}

func (c CustomerService) GetAll()([]entities.Customer,error){

	return c.store.GetAll()
}


func (c CustomerService) Remove(id int) error{

	return c.store.Remove(id)
}

func (c CustomerService) Update(customer entities.Customer,id int) (entities.Customer,error){
	cnt,err:=c.store.Update(customer,id)
	if err!=nil{
		return entities.Customer{},err
	}

	return cnt,nil
}
