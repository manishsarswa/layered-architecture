package store

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"layres_new/entities"
	"log"
	"reflect"
	"testing"
)


func TestCustomerStore_GetByID( t *testing.T){
	var db, err = sql.Open("mysql", "root:Manish@123Sharma@/Customer_services")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)// proper error handling instead of panic in your app
	}
	testcases:=[]struct{
		input int
		output entities.Customer
	}{
		{input: 45, output: entities.Customer{Id: 45, Name: "manish", Dob: "12/12/2000", Address: entities.Address{Id: 28, StreetName: "bikaner", City: "ss", State: "s", CustomerId: 45}}},
		{1234,entities.Customer{}},
	}
	DB:=New(db)
	for i:=range testcases{
		resp, _ := DB.GetByID(testcases[i].input)

		if !reflect.DeepEqual(resp, testcases[i].output) {
			t.Errorf("Failed")
		}
	}
}


func TestCustomerStore_GetByName(t *testing.T) {

	var db, err = sql.Open("mysql", "root:Manish@123Sharma@/Customer_services")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)// proper error handling instead of panic in your app
	}
	testcases:=[]struct{
		input string
		output []entities.Customer
	}{
		{input: "manish", output: []entities.Customer{{Id: 45, Name: "manish", Dob: "12/12/2000", Address: entities.Address{Id: 28, StreetName: "bikaner", City: "ss", State: "s", CustomerId: 45}}}},
		{"1234",[]entities.Customer(nil)},
	}
	DB:=New(db)
	for i:=range testcases{
		resp, _ := DB.GetByName(testcases[i].input)

		if !reflect.DeepEqual(resp, testcases[i].output) {
			t.Errorf("Failed")
		}
	}
}


func TestCustomerStore_Remove(t *testing.T) {

	var db, err = sql.Open("mysql", "root:Manish@123Sharma@/Customer_services")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)// proper error handling instead of panic in your app
	}
	testcases:=[]struct{
		input int
		output entities.Customer
	}{
		{input: 31, output: entities.Customer{}},
		{1234,entities.Customer{}},
	}

		DB:=New(db)
		for i:=range testcases{
			err:= DB.Remove(testcases[i].input)

			if err!=nil {
				t.Errorf("Failed")
			}
		}

}


func TestCustomerStore_Create(t *testing.T) {
	fdb,mock,err:=sqlmock.New()
	if err!=nil{
		log.Fatal("error while opening fdb")
	}
	DB:=New(fdb)
	testcases:=[]struct{
		input entities.Customer
		output entities.Customer
	}{
		{
			entities.Customer{0,"sharma","13/12/2000",entities.Address{0,"sdf","sdfg","ertgfcx",0}},entities.Customer{"46","sharma","13/12/2000",entities.Address{29,"sdf","sdfg","ertgfcx",46}}
		},
	}

}