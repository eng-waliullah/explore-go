package main

import "fmt"


func add(x int, y int) int {
	var res int
	res = x + y
	return res
}


func main(){

	var a int = 10

	var sum = add(a, 6) 

	fmt.Println(sum)

}