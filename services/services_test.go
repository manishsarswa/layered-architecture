package services

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"layres_new/entities"
	"layres_new/store"
	"reflect"
	"testing"
)

func TestCustomerService_GetByID(t *testing.T) {
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

	DB:=New(store.New(db))
	for i:=range testcases{
		resp, _ := DB.GetByID(testcases[i].input)

		if !reflect.DeepEqual(resp, testcases[i].output) {
			t.Errorf("Failed")
		}
	}

}

func TestCustomerService_GetByName(t *testing.T) {
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
	DB:=New(store.New(db))
	for i:=range testcases{
		resp,err:=DB.GetByName(testcases[i].input)
		if err!=nil || !reflect.DeepEqual(resp,testcases[i].output) {
			t.Error("failed")
		}

	}
}

