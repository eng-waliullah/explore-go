// defer fucntion ---> under the hood closer

// parent & closer jodi exists kore tahole closer variable ar address refer share kore

// defer list pointer

// defer --> linked list data stracture use kore

/*

named returned values rules

1. all code executed
2. defer function store kora hobe magic box e
3. return --> all defer fucntion execute korbe
4. return korbe named variables gular values


just return types
1. all code executed
2. defer function store kora hobe magic box e
3. return values are evaluated at this time (store the return value)
4. all defer functions execute korbe


*/

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

	p := func(a int) {
		fmt.Println("ami", a)
	}

	defer p(result)

	defer fmt.Println(result)

	fmt.Println("second ", result)

	defer fmt.Println(5)

	return

}

/*

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

*/

func main() {

	a := calculate()
	fmt.Println("main first", a)

	// b := calc()
	// fmt.Println("main second", b)

}
