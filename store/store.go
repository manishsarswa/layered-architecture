package store

import "C"
import (
	"database/sql"
	"fmt"
	"layres/entities"
)

type CustomerStore struct {
	Db *sql.DB
}

//func (c CustomerStore) fakeDb() (sqlmock.Sqlmock,error,*sql.Db){
//	fDb,mock,err:=sqlmock.New()
//	c.Db=fDb
//	return mock,err,fDb
//}
func (c CustomerStore)CloseDb(){
	c.Db.Close()
}
func New() CustomerStore {
	var Db, err = sql.Open("mysql", "root:Manish@123Sharma@/Customer_services")
	if err != nil {
		panic(err)
	}
	err = Db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	return CustomerStore {Db: Db}
}

func (c CustomerStore) GetCustomerBYId(id int) (entities.Customer, error) {
	rows, err := c.Db.Query("select * from customer inner join address on customer.id=address.cid and customer.id=? ", id)
	if err != nil {
		return entities.Customer{}, err
	}

	var cust entities.Customer

	for rows.Next() {
		rows.Scan(&cust.Id, &cust.Name, &cust.Dob, &cust.Address.Id, &cust.Address.StreetName, &cust.Address.City, &cust.Address.State, &cust.Address.CustomerId)
	}

	return cust, nil
}

func (c CustomerStore) GetCustomerByName(name string) (entities.Customer, error) {
	rows, err := c.Db.Query("select * from customer inner join address on customer.id=address.cid where customer.name=? ", name)
	if err != nil {
		return entities.Customer{}, err
	}
	fmt.Println(name)
	var cust entities.Customer

	for rows.Next() {
		rows.Scan(&cust.Id, &cust.Name, &cust.Dob, &cust.Address.Id, &cust.Address.StreetName, &cust.Address.City, &cust.Address.State, &cust.Address.CustomerId)
	}
	fmt.Println(cust)
	return cust, nil
}

func (c CustomerStore) CreateCustomer(cust entities.Customer) (entities.Customer,error){
	var info[] interface{}
	query:=`insert into customer (name,dob) values(?,?)`
	if cust.Name=="" || cust.Dob==""{
		return entities.Customer{},nil
	}
	if cust.Address.StreetName=="" || cust.Address.City=="" || cust.Address.State==""{
		return entities.Customer{},nil
	}
	info=append(info,&cust.Name)
	info=append(info,&cust.Dob)

	row,_:=c.Db.Exec(query,info...)
	query=`insert into address (street_name,city,state,cid) values(?,?,?,?)`
	var addr[] interface{}

	addr=append(addr,&cust.Address.StreetName)
	addr=append(addr,&cust.Address.City)
	addr=append(addr,&cust.Address.State)

	id,ok1:=row.LastInsertId()
	if ok1!=nil{
		fmt.Println(ok1)
	}
	addr=append(addr,id)
	_,ok:=c.Db.Exec(query,addr...)
	if ok!=nil{
		fmt.Println(ok)
	}
	query=`select * from customer inner join address on customer.id=address.cid where customer.id=?`


	newRow,_:=c.Db.Query(query,id)
	var detail entities.Customer
	for newRow.Next() {
		newRow.Scan(&detail.Id, &detail.Name, &detail.Dob, &detail.Address.Id, &detail.Address.StreetName, &detail.Address.City, &detail.Address.State, &detail.Address.CustomerId)
	}

	return detail,nil
}

func (c CustomerStore) GetCustomer() ([]entities.Customer,error){
	query:=`select * from customer inner join address on customer.id=address.cid`
	rows,ok:=c.Db.Query(query)
	if ok!=nil {
		panic(ok)
	}

	var response []entities.Customer

	defer rows.Close()

	for rows.Next() {
		var detail entities.Customer
		ok = rows.Scan(&detail.Id,&detail.Name,&detail.Dob,&detail.Address.Id,&detail.Address.StreetName,&detail.Address.City,&detail.Address.State,&detail.Address.CustomerId)
		response = append(response, detail)
	}
	return response,nil
}

func (c CustomerStore) RemoveCustomer(id int) error{
	var info[] interface{}
	info=append(info,id)

	query := `delete from customer where id=?`
	_, ok:= c.Db.Exec(query, info...)
	if ok!=nil{
		return ok
	}
	return nil
}

func (c CustomerStore) UpdateCustomer (customer entities.Customer,id int) (entities.Customer,error){
	if customer.Name!=""{
		query:=`update customer set`
		var info [] interface{}
		query+=" name=?"
		info=append(info,customer.Name)
		query+=" where customer.id=?"
		info=append(info,id)
		_,er:=c.Db.Exec(query,info...)

		if er!=nil{
			return entities.Customer{},er
		}
	}

	check:=entities.Address{}
	if  customer.Address!=check {
		query := `update address set `
		var idd []interface{}
		if customer.Address.StreetName != "" {
			idd = append(idd, customer.Address.StreetName)
			query += " street_name=?,"
		}

		if customer.Address.City != "" {
			idd = append(idd, customer.Address.City)
			query += " city=?,"
		}

		if customer.Address.State != "" {
			idd = append(idd, customer.Address.State)
			query += " state=?,"
		}
		query=query[:len(query)-1]
		query += " where address.cid=?"
		idd = append(idd, id)
		_, ok1 := c.Db.Exec(query, idd...)

		if ok1 != nil {
			return entities.Customer{},ok1
		}
	}

	query:=`select * from customer inner join address on customer.id=address.cid where customer.id=?`
	rows,_:=c.Db.Query(query,id)
	var detail entities.Customer
	for rows.Next(){
		rows.Scan(&detail.Id,&detail.Name,&detail.Dob,&detail.Address.Id,&detail.Address.StreetName,&detail.Address.City,&detail.Address.State,&detail.Address.CustomerId)
	}
	if detail.Id==0{
		return entities.Customer{},nil
	}
	return detail,nil
}

