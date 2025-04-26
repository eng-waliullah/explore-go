package main

import "fmt"
import "practice/mathLib"



var a = 20
var b = 30

/*
scope 
	## global
	## local -- if elase -- switch -- function
	## pakage - function or variable must start with capital letter


1. block ---> {} is a block
2. go mod init practice

*/


func main(){
	var p = 30
	var q = 40

	fmt.Println("call in from func/ pakage")

	mathLib.Add(p, q)

	mathLib.adAddd(a, b)

	mathLib.Add(a, p)

}



