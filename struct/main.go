package main

import "fmt"

type User struct {
	Name string // member variable or property
	Age int
}

func main(){
	var user1 User

	user1 = User{ // instantiate
		Name: "Habib",
		Age: 32
	}
}
