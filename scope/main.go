package main

import "fmt"

var msg = "I'm accessible everywhere in main"

func main() {
	x := 10 // Outer scope

	{
		y := 30           // Inner scope
		fmt.Println(x, y) // Valid (10 30)
	}

	// fmt.Println(y) // Compile error: undefined:y

	result := calculate(10, 20)
	fmt.Println(result) // 210

	fmt.Println(msg) // accessible
}

func calculate(a, b int) (result int) { // a, b, and result are function-scoped

	fmt.Println(msg) // accessible
	temp := a * b    // Also function-scoped

	return temp + 10
}
