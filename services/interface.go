package services

import (
	"layres/entities"
	"net/http"
)
type Customer interface {
	GetCustomerBYId(id int)
	GetCustomerByName(name string)
	CreateCustomer(customer entities.Customer)
	GetCustomer(w http.ResponseWriter)
	RemoveCustomer(w http.ResponseWriter,id int)
	UpdateCustomer(customer entities.Customer,id int) entities.Customer
	CloseDb()

}
