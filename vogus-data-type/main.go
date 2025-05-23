package main

import "fmt"

func main() {

	var a int8 = 127
	// var b int8 = -128

	// var c uint8 = 200

	var d float64 = 20.90343453

	var e bool = true

	var s string = "ha"

	ha := "💖"

	var hab rune = '💖'

	fmt.Println(ha)
	fmt.Printf("%c\n", ha)
	fmt.Printf("%d\n", a)
	fmt.Printf("%.2f\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%s\n", s)
	fmt.Printf("%T\n", s)

	fmt.Printf("%c\n", hab)
	fmt.Printf("%T\n", hab)

	// fmt.Println(a, b, c, d)
}
