package main

// func main() {
// 	a := 333

// 	fmt.Println(a)
// }

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0)) // Number of Ps
	for i := 0; i < 10; i++ {
		go func(id int) {
			for {
				fmt.Printf("Goroutine %d running\n", id)
			}
		}(i)
	}
	select {} // Keep main goroutine alive
}
