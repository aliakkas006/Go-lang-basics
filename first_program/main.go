package main

import (
	"fmt"

	"example.com/functionlib"
	"example.com/mathlib"
	"example.com/scopelib"
)

var x = 10

// Internal Memory
const b = 10

var p = 100

func call() {
	add := func(x, y int) {
		z := x + y
		fmt.Println("z = ", z)
	}

	add(5, 6)

	add(p, b)
}

func main() {
	fmt.Println("From mathlib package -")
	sum := mathlib.Add(4, 7)
	fmt.Println(sum)

	fmt.Println("From scopelib package")
	scopelib.Scope()

	shadow()

	fmt.Println(x)

	// Anonymous function
	// IIFE
	func(a, b int) {
		sum := a + b
		println(sum)
	}(3, 5)

	// Function expression or assign function in a variable
	add := func(a, b int) {
		sum := a + b
		println(sum)
	}

	add(1, 2)

	fmt.Println("Call from Higher-order function - ")
	result := functionlib.ProcessOperation(10, 20, add)	
	result(20, 10)	// output: 30

	// Internal Memory
	call()
	fmt.Println("b = ", b)
}

func init() {
	fmt.Println("I will be called first because I am init() -")
	fmt.Println(x)
	x = 20

	// Internal Memory
	fmt.Println("From Init function")
}
