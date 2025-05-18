package main

import "fmt"

/*

1. code segment --> function
2. data segment --> global memory
3. stack --> execution time --> stacl frame
4. heap

## stack (local) memory is faster than global memory
## GC -- garbage collertor

A closur is a function defined within onother functuin and has access to the outer function's variables even after the outer function has finished executing


*/

const a = 10 // constant
var p = 100  // run time

func outer() func() {
	money := 100
	age := 30
	fmt.Println("Age: ", age)

	show := func() {
		money = money + a + p
		fmt.Println(money)
	}
	return show
}

func call() {
	inc1 := outer()
	inc1()
	inc1()

	inc2 := outer()
	inc2()
	inc2()
}

func main() {
	call()
	fmt.Println("call from main func", a)
}

func init() {
	fmt.Println("run  from init func")
}

/*

1. Compile phase (compile time)
2. Execution phase (run time)


## go run main.go --> compile it --> create main file --> then run ./main file automatically
## go build main.go --> compile it --> create main file --> compile it --> create main file manually

***** In Compile phase All Code Allocated In Code Segment // Read Only *****

** In Stack When Function Invoked Then Allocated Memory And This Space Called AS Stack Frame **

** Compile phase in code segment --> global variable in data segment without constant and function -- > init --> main **
** Executio in stack -> stack frame inner stack **

** Code Segment **
 a = 10 // beacause it is constant
 outer = () { ... }
 outerAnonymus1 = () { ... } bind with outer func
 call = () { ... }
 main = () { ... }
 init = () { ... }

 call = () { ... }
 add = () { ... }
 main = () { ... }
 init = () { ... }



*/
