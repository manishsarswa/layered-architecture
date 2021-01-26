package delivery

import (
	"bytes"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"layres_new/entities"
	"layres_new/services"
	"layres_new/store"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHandler_GetByID(t *testing.T) {

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
		output entities.Customer
	}{
		{input: "45", output: entities.Customer{Id: 45, Name: "manish", Dob: "12/12/2000", Address: entities.Address{Id: 28, StreetName: "bikaner", City: "ss", State: "s", CustomerId: 45}}},
		{"1234",entities.Customer{}},
	}

	DB:=New(services.New(store.New(db)))
	for i:=range testcases{
		w:=httptest.NewRecorder()
		req:=httptest.NewRequest(http.MethodGet,"/customer/"+(testcases[i].input),nil)
		req=mux.SetURLVars(req,map[string]string{"id":testcases[i].input})
		DB.GetByID(w,req)
		var customer entities.Customer
		err:=json.Unmarshal(w.Body.Bytes(),&customer)

		if err!=nil{
			t.Log(err)
		}
		if !reflect.DeepEqual(customer, testcases[i].output) {
			t.Errorf("Failed")
		}
	}
}


func TestHandler_GetByName(t *testing.T) {

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

	DB:=New(services.New(store.New(db)))
	for i:=range testcases{
		w:=httptest.NewRecorder()
		req:=httptest.NewRequest(http.MethodGet,"/customer?name="+testcases[i].input,nil)
		req=mux.SetURLVars(req,map[string]string{"name":testcases[i].input})
		DB.GetByName(w,req)
		var customers []entities.Customer
		err:=json.Unmarshal(w.Body.Bytes(),&customers)
		if err!=nil || !reflect.DeepEqual(customers,testcases[i].output){
			t.Error("failed")
		}

	}

}



func TestHandler_Create(t *testing.T) {

	var db, err = sql.Open("mysql", "root:Manish@123Sharma@/Customer_services")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)// proper error handling instead of panic in your app
	}


	testcases:=[]struct{
		input []byte
		output entities.Customer
	}{
		{input: []byte(`{"name": "sdfgh", "dob": "12/12/1985", "address":{"streetName": "wsedrft", "city": "sdfr", "state": "qwedrf"}}`), output: entities.Customer{Id:56, Name: "sdfgh", Dob: "12/12/1985", Address: entities.Address{Id:39, StreetName: "wsedrft", City: "sdfr", State: "qwedrf",CustomerId: 56}}},
		{input: []byte(`{"name": "sdfgh", "dob": "12/12/2010", "address":{"streetName": "wsedrft", "city": "sdfr", "state": "qwedrf"}}`), output: entities.Customer{}},
	}


	DB:=New(services.New(store.New(db)))
	for i:=range testcases{
		w:=httptest.NewRecorder()
		req:=httptest.NewRequest(http.MethodGet,"/customer",bytes.NewBuffer(testcases[i].input))
		DB.Create(w,req)
		var customers entities.Customer
		err:=json.Unmarshal(w.Body.Bytes(),&customers)
		if customers.Id!=0 {
			if err != nil || !reflect.DeepEqual(customers, testcases[i].output) {
				t.Error("failed")
			}
		}

	}


}
