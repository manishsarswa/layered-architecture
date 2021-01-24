package main

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(result)
		if err!=nil{
			t.Log(err)
		}


		if !reflect.DeepEqual(result,testcases[i].Output){
			t.Errorf("Expected result  is %v and got result is %v",result,testcases[i].Output)
		}

	}


}

