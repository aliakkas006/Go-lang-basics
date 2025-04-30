package main

import "fmt"

/*
Variadic function:
  - A variadic function in Go is a function that accepts a variable number of arguments of the same type.
  - use the ... (ellipsis) syntax before the type to define it.
  - example: func funcName(arg ...type)
*/
func variadicFunc(numbers ...int) { // numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers, len(numbers), cap(numbers)) // [1, 2, 3, 4, 5], len = 5, cap = 5
}

func sum(nums ...int) int { // nums := []int{1, 2, 3}
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

/* 
    * When we call a variadic function, Go converts the arguments into a slice
    * The variadic parameter is implemented as a slice under the hood
    * The compiler generates code to automatically create this slice

    * Calling with no variadic args creates a nil slice.
        - numbers := []int
*/

func main() {
	variadicFunc(1, 2, 3, 4, 5)

	summation2 := sum(1, 2, 3)      // 6
	summation := sum(1, 2, 3, 4, 5) // 15

	fmt.Println(summation)
	fmt.Println(summation2)
}
