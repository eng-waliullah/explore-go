package main

import "fmt"

var arr3 = [3]string{"Ami", "Tanvirer", "Mama"}

func main() {

	// var arr [2]int

	// arr2 := [2]int {2, 7}

	// arr[1] = 6

	// arr3[0] = "Hello"
	// arr3[2] = "Mami"

	// fmt.Println(arr)
	// fmt.Println(arr2)
	// fmt.Println(arr3)

	// numbers := [8]int{1, 2, 3, 4, 5, 6, 7, 8}

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	// for index, value := range numbers {
	// 	fmt.Println(index, value)
	// }

	// var multi [3][3]int

	// multi := [3][3]int{
	// 	{1, 2, 3},
	// 	{1, 2, 3},
	// 	{1, 2, 3},
	// }
	// fmt.Println(multi)

	// arr1 := [5]int{1, 2, 3, 4, 5}
	// arr2 := [5]int{1, 2, 3, 4, 5}

	// fmt.Println(arr1)
	// fmt.Println(arr2)

	var arr3 [5]int

	for i := 0; i < len(arr3); i++ {
		fmt.Scan(&arr3[i])
	}
	fmt.Println(arr3)

}
