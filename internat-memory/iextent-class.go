package main

import "fmt"



/*

1. code segment --> function
2. data segment --> global memory
3. stack --> execution time --> stacl frame
4. heap

## stack (local) memory is faster than global memory
## GC -- garbage collertor

*/

 

const a = 10
var p = 100


func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println(z)
	}

	add(5, 6)
	add(p, a)
}


func main(){
	call()
	fmt.Println(a)
}


func init(){
	fmt.Println("from init func")
}



/*

1. Compile phase
2. Execution phase


## go run main.go --> compile it --> create main file --> then run ./main file automatically
## go build main.go --> compile it --> create main file --> compile it --> create main file manually

***** In Compile phase All Code Allocated In Code Segment // Read Only *****

** In Stack When Function Invoked Then Allocated Memory And This Space Called AS Stack Frame **

** Compile phase in code segment --> global variable in data segment without constant and function -- > init --> main **
** Executio in stack -> stack frame inner stack **

** Code Segment **
 a = 10 // beacause it is constant
 call = () { ... }
 add = () { ... }
 main = () { ... }
 init = () { ... }



*/