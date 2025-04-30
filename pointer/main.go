package main

import "fmt"


func print(nums *[3]int) {
	fmt.Println(nums)
}

type User struct {
	Name string
	Age int
	Salary float64
	FavFoods []string
}


func main(){

	// pointer // address // memory ( ram / hard disk )

	// x := 10

	// p := &x // ampersane --> address of 

	// *p = 1000

	// fmt.Println("X:", x)
	// fmt.Println("Address:", p) // address of x
	// fmt.Println("value at the address:", *p) // *  value of address


	// pass by value
	// pass by reference

	habib := User {
		Name: "Habib",
		Age: 27,
		Salary: 0,
	}

	fmt.Println(habib)

	p := &habib

	fmt.Println(p.Name)
	


	// arr := [3]int {1, 2, 3}

	// print(&arr)

 
}


