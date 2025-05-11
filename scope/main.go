package main

import "fmt"

var msg = "Package-scoped variable" // Not exported
var Msg = "Package-scoped variable" // Exported

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

func calculate(a, b int) (result int) { // a, b, and result are function-scoped (Not Exported)

	fmt.Println(msg) // accessible
	temp := a * b    // Also function-scoped

	return temp + 10
}

func Calculate(a, b int) (result int) { // a, b, and result are function-scoped (Exported)

	fmt.Println(msg) // accessible
	temp := a * b    // Also function-scoped

	return temp + 10
}

var (
	name = "ali"
	age  = 30
)
