package store

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"layres_new/entities"
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

//
//func TestCustomerStore_GetByName(t *testing.T) {
//
//	var db, err = sql.Open("mysql", "root:Manish@123Sharma@/Customer_services")
//	if err != nil {
//		panic(err)
//	}
//	err = db.Ping()
//	if err != nil {
//		panic(err)// proper error handling instead of panic in your app
//	}
//	testcases:=[]struct{
//		input string
//		output []entities.Customer
//	}{
//		{input: "manish", output: []entities.Customer{{Id: 45, Name: "manish", Dob: "12/12/2000", Address: entities.Address{Id: 28, StreetName: "bikaner", City: "ss", State: "s", CustomerId: 45}}}},
//		{"1234",[]entities.Customer(nil)},
//	}
//	DB:=New(db)
//	for i:=range testcases{
//		resp, _ := DB.GetByName(testcases[i].input)
//
//		if !reflect.DeepEqual(resp, testcases[i].output) {
//			t.Errorf("Failed")
//		}
//	}
//}

func TestCustomerStore_GetByName(t *testing.T) {

	fdb,mock,err:=sqlmock.New()
	if err != nil {
		t.Fatalf("Cannot connect to Mock DataBase")
	}
	testcases:=[]struct{
		input string
		err error
		output []entities.Customer

 	}{
		{input: "manish", err:nil ,output: []entities.Customer{{Id: 45, Name: "manish", Dob: "12/12/2000", Address: entities.Address{Id: 28, StreetName: "bikaner", City: "ss", State: "s", CustomerId: 45}}}},
		{"1234",errors.New("Name does not exist"),[]entities.Customer(nil)},
	}
	DB:=New(fdb)
	str:=[]string{"id","name","dob","id","street_name","city","state","cid"}
	for i,val:=range testcases{
		var row *sqlmock.Rows
		row=sqlmock.NewRows(nil).AddRow()
		if val.output!=nil {
			row = sqlmock.NewRows(str).AddRow(val.output[i].Id, val.output[i].Name, val.output[i].Dob, val.output[i].Address.Id, val.output[i].Address.StreetName, val.output[i].Address.City, val.output[i].Address.State, val.output[i].Address.CustomerId)
		}
		mock.ExpectQuery("select *").WithArgs(val.input).WillReturnRows(row).WillReturnError(val.err)
		result,err:=DB.GetByName(val.input)
		if err!=nil{
			if err!=val.err{
				t.Errorf("Failed")
			}
		}
		if !reflect.DeepEqual(result,testcases[i].output) {
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

//
//func TestCustomerStore_Create(t *testing.T) {
//	fdb,mock,err:=sqlmock.New()
//	if err!=nil{
//		log.Fatal("error while opening fdb")
//	}
//	DB:=New(fdb)
//	testcases:=[]struct{
//		input entities.Customer
//		output entities.Customer
//	}{
//
//		{input: entities.Customer{Name: "sharma", Dob: "13/12/2000", Address: entities.Address{StreetName: "sdf", City: "sdfg", State: "ertgfcx"}}, output: entities.Customer{Id: 46, Name: "sharma", Dob: "13/12/2000", Address: entities.Address{Id: 29, StreetName: "sdf", City: "sdfg", State: "ertgfcx", CustomerId: 46}}},
//	}
//}
