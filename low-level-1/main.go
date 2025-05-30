package main

import "fmt"

func foo(b int) int {
	b = b + 1
	return b
}

func cat(c int) int {
	d := foo(c)
	return d
}

func main() {
	a := cat(10)
	fmt.Println(a)
}
