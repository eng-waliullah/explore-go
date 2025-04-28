package main

import "fmt"

var a = 20
var b = 30

/*

1. code segment --> function
2. data segment --> global memory
3. stack --> execution time --> stacl frame
4. heap

## stack (local) memory is faster than global memory
## GC -- garbage collertor

*/



func add(x int, y int){
 	var z = x + y

	fmt.Println(z)
}

func main(){
	// var p = 30
	// var q = 40

	add(a, b)
	add(a, 10)

}

func init(){
	fmt.Println("from init func")
}