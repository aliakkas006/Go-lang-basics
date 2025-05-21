package main

import "fmt"

var arr_3 = [2]string{"Ali", "Akkas"}

func main() {
	var arr [3]int

	arr[0] = 1
	arr[1] = 2

	fmt.Println(arr)

	// Short-hand declaration:
	arr_2 := [3]int{1, 2, 3}

	fmt.Println(arr_2)
	fmt.Println(arr_3)

	// Multidimensional array
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(matrix)
}
