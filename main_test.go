package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"layres/delivery"
	"layres/entities"
	"layres/services"
	"layres/store"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Define struct{
	Input string
	Output entities.Customer
}


func TestGetByName(t *testing.T) {

	testcases:=[]struct{
		input string
		output entities.Customer
	}{
		{input: "?name=manish", output: entities.Customer{Id: 31, Name: "manish", Dob: "12/12/2000", Add: entities.Address{Id: 19, StreetName: "bikaner", City: "s", State: "ss", CustomerId: 31}},},
		{"?name=pankaj",entities.Customer{}},
	}
	datastore := store.New()
	defer datastore.CloseDb()
	service := services.New(datastore)
	handler := delivery.New(service)
	for i:= range testcases{
		w:=httptest.NewRecorder()
		req:=httptest.NewRequest(http.MethodGet,"/customer"+testcases[i].input,nil)
		req=mux.SetURLVars(req,map[string]string{"id":testcases[i].input})

		handler.GetCustomerByName(w,req)
		var c entities.Customer
		err:=json.Unmarshal(w.Body.Bytes(),&c)
		if err!=nil{
			t.Log(err)
		}

		if !reflect.DeepEqual(c,testcases[i].output){
			t.Errorf("Failed in %v testcase",i)
		}
	}
}


func TestGetCustomerById(t *testing.T){
	testcases:=[]Define{
		{Input: "31", Output: entities.Customer{Id: 31, Name: "manish", Dob: "12/12/2000", Add: entities.Address{Id: 19, StreetName: "bikaner", City: "s", State: "ss", CustomerId: 31}}},
		{"1234",entities.Customer{}},
	}
	datastore := store.New()
	defer datastore.CloseDb()
	service := services.New(datastore)
	handler := delivery.New(service)
	for i:=range testcases{
		req:=httptest.NewRequest(http.MethodGet,"/customer/"+testcases[i].Input,nil)
		req=mux.SetURLVars(req,map[string]string{"id":testcases[i].Input})
		w:=httptest.NewRecorder()
		handler.GetCustomerById(w,req)
		var result entities.Customer
		err:=json.Unmarshal(w.Body.Bytes(),&result)
		if err!=nil{
			t.Log(err)
		}


		if !reflect.DeepEqual(result,testcases[i].Output){
			t.Errorf("Expected result  is %v and got result is %v",result,testcases[i].Output)
		}

	}
}

//func TestCreateCustomer(t *testing.T) {
//
//
//	fdb,mock,err:=sqlmock.New()
//
//	if err!=nil{
//		fmt.Println(err)
//	}
//	testcases:=[]struct{
//		input entities.Customer
//		output entities.Customer
//	}{
//		{input: entities.Customer{Name: "Sharma", Dob: "13/10/2000", Add: entities.Address{StreetName: "7th cross", City: "Bangalore", State: "karnataka"}}, output: entities.Customer{Id: 1, Name: "Sharma", Dob: "13/10/2000", Add: entities.Address{Id: 1, StreetName: "7th cross", City: "Bangalore", State: "karnataka", CustomerId: 1}}},
//	}
//
//	for i:=range testcases{
//		mock.NewRows([]string{"id","Name","Dob"}).AddRow(testcases[i].output.Id,testcases[i].output.Name,testcases[i].output.Dob)
//		mock.ExpectExec("insert into customer (name,dob) values*").WithArgs(testcases[i].input.Name,testcases[i].input.Dob).WillReturnResult(sqlmock.NewResult(int64(testcases[i].output.Id),1))
//
//		mock.NewRows([]string{"id","StreetName","City","State","CustomerId"}).AddRow(testcases[i].output.Id,testcases[i].output.Add.StreetName,testcases[i].output.Add.City,testcases[i].output.Add.State,testcases[i].output.Add.CustomerId)
//		mock.ExpectExec("insert into address (street_name,city,state,cid) values*").WithArgs(testcases[i].input.Add.StreetName,testcases[i].input.Add.City,testcases[i].input.Add.State,testcases[i].output.Add.CustomerId).WillReturnResult(sqlmock.NewResult(int64(testcases[i].output.Id),1))
//		body,_:=json.Marshal(testcases[i].input)
//		req:=httptest.NewRequest(http.MethodPost,"/customer/",bytes.NewBuffer(body))
//		w:=httptest.NewRecorder()
//
//
//
//	}
//}
//


func TestCreateCustomer(t *testing.T){
	testcases:=[]struct{
		input []byte
		output entities.Customer
	}{
		{input: []byte(`{"Name":"sharma","Dob":"13/12/2000","StreetName":"5th cross","City":"bangalore","State":"Karnataka"}`), output: entities.Customer{Id: 34, Name: "sharma", Dob: "13/12/2000", Add: entities.Address{Id: 22, StreetName: "5th cross", City: "bangalore", State: "karnataka", CustomerId: 34}}},
		{input: []byte(`{"Name":"Ankit","Dob":"14/10/2000","StreetName":"5th cross","City":"bangalore","State":"Karnataka"}`), output: entities.Customer{Id: 35, Name: "Ankit", Dob: "14/10/2000", Add: entities.Address{Id: 23, StreetName: "5th cross", City: "bangalore", State: "karnataka", CustomerId: 35}}},

	}

	datastore := store.New()
	defer datastore.CloseDb()
	service := services.New(datastore)
	handler := delivery.New(service)

	for i:=range testcases{
		req:=httptest.NewRequest(http.MethodPost,"/customer/",bytes.NewBuffer(testcases[i].input))
		w:=httptest.NewRecorder()
		handler.CreateCustomer(w,req)
		var c entities.Customer
		json.Unmarshal(w.Body.Bytes(),&c)
		if !reflect.DeepEqual(c,testcases[i].output){
			t.Errorf("Failed in %v testcase",i)
		}

	}
}

func TestUpdateCustomer(t *testing.T) {
	testcases := []struct {
		input  string
		body   []byte
		output entities.Customer
	}{
		{"31", []byte(`{"Name":"manishsharma"}`), entities.Customer{31, "manishsharma", "12/12/2000", entities.Address{Id: 19, StreetName: "bikaner", City: "s", State: "ss", CustomerId: 31}}},
	}

	datastore := store.New()
	defer datastore.CloseDb()
	service := services.New(datastore)
	handler := delivery.New(service)

	for i :=range testcases{

		w:=httptest.NewRecorder()
		req:=httptest.NewRequest(http.MethodPut,"/customer?id="+string(testcases[i].input),bytes.NewBuffer(testcases[i].body))
		handler.UpdateCustomer(w,req)
		var c entities.Customer
		json.Unmarshal(w.Body.Bytes(),&c)
		if !reflect.DeepEqual(c,testcases[i].output){
			t.Errorf("Failed in %v testcase ",i)
		}
	}
}

func TestRemoveCustomer(t *testing.T){
	testcases:=[]Define{
		{Input: "31", Output: entities.Customer{Id: 31, Name: "manish", Dob: "12/12/2000", Add: entities.Address{Id: 19, StreetName: "bikaner", City: "s", State: "ss", CustomerId: 31}}},
		{"0",entities.Customer(nil)},
		{"1234",entities.Customer(nil)},
	}

	datastore := store.New()
	defer datastore.CloseDb()
	service := services.New(datastore)
	handler := delivery.New(service)
	for i :=range testcases{
		w:=httptest.NewRecorder()
		req:=httptest.NewRequest(http.MethodDelete,"/customer/"+testcases[i].Input,nil)
		req=mux.SetURLVars(req,map[string]string{"id":testcases[i].Input})
		handler.RemoveCustomer(w,req)
		var c entities.Customer
		json.Unmarshal(w.Body.Bytes(),&c)
		if !reflect.DeepEqual(c,testcases[i].Output){
			t.Errorf("Failed in %v testcase ",i)
		}

	}
}