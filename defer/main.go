package main

import "fmt"

/*

func a() {
	i := 0

	fmt.Println("firsts", i) // 0

	defer fmt.Println("second", i) // 0  --> fmt.Println("second", 0)

	i++

	defer fmt.Println("fourth", i)

	fmt.Println("third", i) // 1

	return
}


*/

func calculate() (result int) {

	fmt.Println("first ", result)

	defer func() {
		result = result + 10

		fmt.Println("defer ", result)
	}()

	result = 5

	fmt.Println("second ", result)

	return

}

func calc() int {

	result := 0

	fmt.Println("first ", result)

	defer func() {
		result = result + 10

		fmt.Println("defer ", result)
	}()

	result = 5

	fmt.Println("second ", result)

	return result

}

func main() {

	a := calculate()
	b := calc()

	fmt.Println(a)
	fmt.Println(b)

}
