package store

import (
	"database/sql"
	"layres_new/entities"
	"layres_new/errors"
)

type CustomerStore struct {
	DB *sql.DB
}


func (c CustomerStore) CloseDB(){
	c.DB.Close()
}
func New(db *sql.DB) Customer {
	return CustomerStore {DB: db}
}

func (c CustomerStore) GetByID(id int) (entities.Customer, error) {
	row, err := c.DB.Query("select * from customer inner join address on customer.id=address.cid and customer.id=? ", id)
	if err != nil {
		return entities.Customer{}, errors.DBError
	}

	var customer entities.Customer

	for row.Next() {
		row.Scan(&customer.Id, &customer.Name, &customer.Dob, &customer.Address.Id, &customer.Address.StreetName, &customer.Address.City, &customer.Address.State, &customer.Address.CustomerId)
	}

	return customer, nil
}

func (c CustomerStore) GetByName(name string) ([]entities.Customer, error) {
	rows, err := c.DB.Query("select * from customer inner join address on customer.id=address.cid where customer.name=? ", name)
	if err != nil {
		return []entities.Customer(nil), errors.DBError
	}

	var customers []entities.Customer

	for rows.Next() {
		var customer entities.Customer
		rows.Scan(&customer.Id, &customer.Name, &customer.Dob, &customer.Address.Id, &customer.Address.StreetName, &customer.Address.City, &customer.Address.State, &customer.Address.CustomerId)
		customers=append(customers,customer)
	}
	return customers, nil
}

func (c CustomerStore) Create(customer entities.Customer) (entities.Customer,error){
	if customer.Name=="" || customer.Dob==""{
		return entities.Customer{},errors.BadRequest
	}

	if customer.Address.StreetName=="" || customer.Address.City=="" || customer.Address.State==""{
		return entities.Customer{},errors.BadRequest
	}

	query:=`insert into customer (name,dob) values(?,?)`

	result,_:=c.DB.Exec(query,customer.Name,customer.Dob)
	query=`insert into address (street_name,city,state,cid) values(?,?,?,?)`

	id,err:=result.LastInsertId()
	if err!=nil{
		return entities.Customer{},errors.BadRequest
	}
	_,err=c.DB.Exec(query,customer.Address.StreetName,customer.Address.City,customer.Address.State,id)
	if err!=nil{
		return entities.Customer{},errors.BadRequest
	}
	query=`select * from customer inner join address on customer.id=address.cid where customer.id=?`

	row,_:=c.DB.Query(query,id)
	var customer1 entities.Customer
	for row.Next() {
		row.Scan(&customer1.Id, &customer1.Name, &customer1.Dob, &customer1.Address.Id, &customer1.Address.StreetName, &customer1.Address.City, &customer1.Address.State, &customer1.Address.CustomerId)
	}

	return customer1,nil
}

func (c CustomerStore) GetAll() ([]entities.Customer,error){
	query:=`select * from customer inner join address on customer.id=address.cid`

	rows,err:=c.DB.Query(query)
	if err!=nil {
		return []entities.Customer(nil),errors.DBError
	}

	var customers []entities.Customer
	defer rows.Close()

	for rows.Next() {
		var customer entities.Customer
		rows.Scan(&customer.Id,&customer.Name,&customer.Dob,&customer.Address.Id,&customer.Address.StreetName,&customer.Address.City,&customer.Address.State,&customer.Address.CustomerId)
		customers = append(customers, customer)
	}
	return customers,nil
}

func (c CustomerStore) Remove(id int) error{

	query := `delete from customer where id=?`
	_, err:= c.DB.Exec(query, id)
	if err!=nil{
		return errors.NotFound
	}
	return nil
}

func (c CustomerStore) Update(customer entities.Customer,id int) (entities.Customer,error){
	var info []interface{}
	if customer.Name!="" {
		query:=`update customer set`
		query+=" name=?"
		info=append(info,customer.Name)
		query+=" where customer.id=?"
		info=append(info,id)
		_,er:=c.DB.Exec(query,info...)

		if er!=nil{
			return entities.Customer{},errors.BadRequest
		}
	}
	var info1 []interface{}
	check:=entities.Address{}
	if  customer.Address!=check {
		query := `update address set `

		if customer.Address.StreetName != "" {
			info1 = append(info1, customer.Address.StreetName)
			query += " street_name=?,"
		}

		if customer.Address.City != "" {
			info1 = append(info1, customer.Address.City)
			query += " city=?,"
		}

		if customer.Address.State != "" {
			info1 = append(info1, customer.Address.State)
			query += " state=?,"
		}
		query=query[:len(query)-1]
		query += " where address.cid=?"
		info1 = append(info1, id)
		_, err := c.DB.Exec(query, info1...)


		if err != nil {
			return entities.Customer{},errors.BadRequest
		}
	}

	query:=`select * from customer inner join address on customer.id=address.cid where customer.id=?`
	row,_:=c.DB.Query(query,id)
	var customer1 entities.Customer
	for row.Next(){
		row.Scan(&customer1.Id,&customer1.Name,&customer1.Dob,&customer1.Address.Id,&customer1.Address.StreetName,&customer1.Address.City,&customer1.Address.State,&customer1.Address.CustomerId)
	}
	if customer1.Id==0{
		return entities.Customer{},errors.BadRequest
	}
	return customer1,nil
}

