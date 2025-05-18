package main

import "fmt"

/*


func counterUpdater() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

*/

func multiplier(factor int) func(n int) int {
	return func(n int) int {
		return factor * n
	}
}

func main() {

	dpoble := multiplier(2)
	triple := multiplier(3)

	fmt.Println(dpoble(2))
	fmt.Println(triple(3))

	/*

		count1 := counterUpdater()
		count2 := counterUpdater()

		fmt.Println(count1())
		fmt.Println(count1())
		fmt.Println(count2())
		fmt.Println(count2())
		fmt.Println(count2())
		fmt.Println(count1())

	*/
}
