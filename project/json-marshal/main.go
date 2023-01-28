package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Salons struct {
	SalonKey string `json:"salon_key"`
	EmployeeKey string `json:"employee_key"`
	CustomerKey string `json:"customer_key"`
	EventKey string `json:"event_key"`
}

func main(){

	bArr := readFile("salons.json")
	salon := new(Salons)

	json.Unmarshal(bArr,salon)

	// serialize
	fmt.Println(salon)
	fmt.Println(salon.SalonKey)
	fmt.Println(salon.EmployeeKey)
	fmt.Println(salon.CustomerKey)

	// deserialize
	desSalon, _ := json.Marshal(salon)
	fmt.Println(desSalon) // []byte
	fmt.Println(string(desSalon))
}

func readFile(f string) []byte {	
	
	d ,err := ioutil.ReadFile(fmt.Sprintf("public/%s",f))
	if err != nil {
		panic(err)
	}

	return d
}