package main

import (
	"fmt"
	"time"
)

var a = 10

const p = 11

func printHello(num int) {
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

}

/*


build main file will save in HDD

threed stack re execute kore

OS kurnel aise stack re execute kore



=========== Code Segment ==========

i. p = 10
ii. printHello = printHello(num int) { .... }
iii. main = main() { .... }




*/
