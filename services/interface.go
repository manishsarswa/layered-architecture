package services

import (
	"layres_new/entities"
)
type Customer interface {
	GetByID(id int)  (entities.Customer,error)
	GetByName(name string)  ([]entities.Customer,error)
	Create(customer entities.Customer)  entities.Customer
	GetAll() ([]entities.Customer,error)
	Remove(id int)  error
	Update(customer entities.Customer,id int)  entities.Customer
	CloseDB()
}
