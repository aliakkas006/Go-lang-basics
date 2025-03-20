package functionlib

import "fmt"

/**
 * Function is a block of code that performs a specific task.
 * A parameter is a variable named in the function definition
 * An argument is the actual value that is passed to the function when it is called.
 */

/**
 * Function types:
 * 1. Standard function (named function)
 * 2. Init function ->
 		- We Can't call this function, Computer calls this function autometically
		- It will called at the beginning of the program execution (even before main function's called)
 * 3. Anonymous function ->
 		- An anonymous function is a function that doesn’t have a name.
 		- It is useful when you want to create an inline function.
 		- In Go, an anonymous function can also form a closure.
 		- An anonymous function is also known as a function literal.
		- We Can assign an anonymous function to a variable. This variable can then be called like a regular function.
		- Anonymous functions can accept arguments.
 * 4. IIFE -> Immediately Invoked Function Expression
*/

/**
* Higher-order function (HOF or First-class function):
	- A function that takes another function as an argument or returns a function as a result is called a higher-order function.
	- In Go, functions are first-class citizens, which means that functions can be assigned to variables, passed as arguments to other functions, and returned from other functions.

* Any of the following 3 conditions must be met for a function to be considered a higher-order function:
	1. Function that takes a function as a parameter (parameter => function)
	2. Function that returns a function	(return => function)
	3. both

* Callback function:
	- A callback function is a function that is passed as an argument to another function.
	- The callback function is called inside the other function.

* First class citizen:
	- Assigned data type to a variable 					
*/

// Higher-order function
func ProcessOperation (a, b int, cb func (x, y int)) func(x, y int) {
	// Execute op func
	cb(a, b)

	// Return div func
	return div
}

func div (x, y int) {
	res := x / y
	fmt.Println(res)
}
