package main

import "fmt"

func print(numbers *[5]int) { // pointer to an array of 5 integers
	fmt.Println(numbers) // &[1 2 3 4 5]
}

func main() {
	x := 10 // Only value hold | 10

	ptr := &x // Address of x | 0xc00000a0e8 | 824633759720

	val := *ptr // Value at address ptr | 10

	*ptr = 20 // Re-assign the value of x at the address of ptr | 20

	fmt.Println("Address of x = ", ptr)
	fmt.Println("Value at address ptr = ", val)
	fmt.Println("Final Value of x = ", x)

	// Array print with the help of pointers
	arr := [5]int{1, 2, 3, 4, 5}

	print(&arr) // pass by reference
}
