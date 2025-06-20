package main

/*

------ main go routine, stack create korbe heap e.


----------- Go Routine ~~~ only 2 kb (thread hote hole stack lagbe ar stack 8mb)
----------- or mini thread
----------- or virtual thread
----------- or logical thread




1. Go Routine Scheduler
2. Heap Allocator (K initialize kore )
3. Garbage Collertor
4. Logical Processors Create kore
5.


## code execute kore akmatro CPU



*/

import (
	"fmt"
	"time"
)

var a = 10

const p = 11

func printHello(num int) {
	time.Sleep(5 * time.Second)
	fmt.Println("Hello Habiba", "--", num)
}

func main() {

	go printHello(1)
	go printHello(2)
	go printHello(3)
	go printHello(4)
	go printHello(5)

	println(a, " ", p)

	time.Sleep(5 * time.Second)
	fmt.Println("Hello Habiba")
}

/*


a. build main file will save in HDD
b. threed stack re execute kore
c. OS kurnel aise stack re execute kore



=========== Code Segment ==========

i. p = 10
ii. printHello = printHello(num int) { .... }
iii. main = main() { .... }




*/
