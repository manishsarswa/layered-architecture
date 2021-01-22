package entities


type Address struct {
Id int          `json:id`
StreetName string `json:streetName`
City string			`json:city`
State string        `json:state`
CustomerId int      `json:customerId`
}

type Customer struct {
	Id   int    `json:id`
	Name string `json:name`
	Dob  string `json:dob`
	Add  Address
}