package main

import (
	"encoding/json"
	"fmt"
	
)

type User struct {
	ID    int `json:"-"`
	Name  string`json:"name"`
	Email string `json:"email"`
}

type Product struct {
	ID    int `transform:"lower"`
	Name  string`transform:"lower"`
}	

func main() {

	u := User{
		ID:    1,
		Name:  "Ortiz",
		Email: "a@gmail.com",
	}

	ujson, _ := json.MarshalIndent(u, "", " ")
	fmt.Println(string(ujson))

	p := Product{
		ID:    1,
		Name:  "Phone",
	}
	Tr(&p)
	fmt.Println(p)
}