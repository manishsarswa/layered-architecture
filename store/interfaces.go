package store

import "layres/entities"

type Customer interface {
	GetCustomerBYId(id int) (entities.Customer, error)
	GetCustomerByName(name string) (entities.Customer,error)
	CreateCustomer(customer entities.Customer) (entities.Customer,error)
	CloseDb()
	GetCustomer() ([]entities.Customer,error)
	RemoveCustomer(id int) error
	UpdateCustomer(customer entities.Customer,id int) (entities.Customer,error)
}