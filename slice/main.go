package main

import "fmt"

func changeSlice(p []int) []int {
	p[0] = 10
	p = append(p, 11)
	return p
}

func print(nums ...int) { // variadic function as like spreed operator in JS
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(cap(nums))
}

func main() {

	// pointer // address // memory ( ram / hard disk )

	/*
		arr := [5]string{"This", "is", "GO", "programmig", "language"}

		fmt.Println(arr)

		s := arr[1:4]
		fmt.Println(s)
		fmt.Println(len(s))

		s1 := s[1:2]
		fmt.Println(s1)
		fmt.Println(len(s1))
		fmt.Println(cap(s1))

		s2 := []int{1, 2, 3} // slice literal
		fmt.Println("slice: ", s2, "len: ", len(s2), "cap: ", cap(s2))

		s3 := make([]int, 3)
		s3[0] = 5
		fmt.Println("slice: ", s3, "len: ", len(s3), "cap: ", cap(s3))

		s4 := make([]int, 3, 5)
		s4[0] = 5
		// s4[3] = 7
		fmt.Println("slice: ", s4, "len: ", len(s4), "cap: ", cap(s4))

		var g []int //  empty or nil slice
		g = append(g, 1)
		fmt.Println("slice: ", g, "len: ", len(g), "cap: ", cap(g))

	*/

	/*

		var x []int      // [], len = 0, cap = 0
		x = append(x, 1) // [1], len = 1, cap = 1
		x = append(x, 2) // [1, 2], len = 2, cap = 2
		x = append(x, 3) // [1, 2, 3], len = 3, cap = 3

		y := x

		x = append(x, 4) // [1, 2, 3, 4], len = 4, cap = 4
		y = append(y, 5) // [], len = 0, cap = 0

		x[0] = 10

		fmt.Println("slice: ", x, "len: ", len(x), "cap: ", cap(x)) // [10 2 3 5]
		fmt.Println("slice: ", y, "len: ", len(y), "cap: ", cap(y)) // [10 2 3 5]
	*/

	p := []int{1, 2, 3, 4, 5}

	p = append(p, 6)
	p = append(p, 7)

	a := p[4:]

	q := changeSlice(a)

	fmt.Println(p) //  [1 2 3 4 10 6 7]
	fmt.Println(q) //  [10 6 7 11]

	fmt.Println(p[0:8]) //  [1 2 3 4 10 6 7 11]

	// ab := []int{1, 2, 3, 3}
	// ab := make([]int, 6, 10)
	// var ab []int
	// ab = append(ab, 12)

	// ab := []int{1, 2, 3}
	// ab = append(ab, 4)
	// ab = append(ab, 5)

	// // ab :=

	// fmt.Println(ab)
	// fmt.Println(len(ab))
	// fmt.Println(cap(ab))
	/*
		var x []int
		x = append(x, 1)
		x = append(x, 2)
		x = append(x, 3)

		y := x
		x = append(x, 4)
		y = append(y, 5)

		x[0] = 10
		fmt.Println(x)
		fmt.Println(y)
	*/

}

/*

1. slice from existing array
2. slice from slice a array
3. slice literal
4. make function with length
5. make function with length and capacity
5. empty or nil slice
6. slice array underlying array rule --> 1024 -> 100% then 25% increase


*/
