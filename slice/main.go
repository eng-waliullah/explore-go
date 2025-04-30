package main

import "fmt"


// func print(nums *[3]int) {
// 	fmt.Println(nums)
// }

// type User struct {
// 	Name string
// 	Age int
// 	Salary float64
// 	FavFoods []string
// }


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

	// habib := User {
	// 	Name: "Habib",
	// 	Age: 27,
	// 	Salary: 0,
	// }

	// fmt.Println(habib)

	// p := &habib

	// fmt.Println(p.Name)
	


	// arr := [5]string { "This", "is", "GO", "programmig", "language" }

	// fmt.Println(arr)

	// s := arr[1: 4]
	// fmt.Println(s)
	// fmt.Println(len(s))

	// s1 := s[1: 2]
	// fmt.Println(s1)
	// fmt.Println(len(s1))
	// fmt.Println(cap(s1))

 


	// s := [] int { 1, 2, 3 } // slice literal
	// fmt.Println("slice: ", s, "len: ", len(s), "cap: ", cap(s))


	// s := make([]int, 3)
	// s[0] = 5
	// fmt.Println("slice: ", s, "len: ", len(s), "cap: ", cap(s))




	// s := make([]int, 3, 5)
	// s[0] = 5
	// s[3] = 7
	// fmt.Println("slice: ", s, "len: ", len(s), "cap: ", cap(s))

	// var x []int //  empty or nil slice
	// x = append(x, 1)
	// fmt.Println("slice: ", x, "len: ", len(x), "cap: ", cap(x))



	var x []int 		// [], len = 0, cap = 0
	x = append(x, 1)	// [1], len = 1, cap = 1
	x = append(x, 2)	// [1, 0], len = 2, cap = 2
	x = append(x, 3)	// [], len = 0, cap = 0
	x = append(x, 30)	

	// y := x

	// x = append(x, 4)	// [], len = 0, cap = 0
	// y = append(y, 5)	// [], len = 0, cap = 0

	// x[0] = 10


	fmt.Println("slice: ", x, "len: ", len(x), "cap: ", cap(x))
	// fmt.Println("slice: ", y, "len: ", len(y), "cap: ", cap(y))



}



/*

1. slice from existing array
2. slice from slice a array
3. slice literal
4. make function with length
5. make function with length and capacity
5. empty or nil slice


*/