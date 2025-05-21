package main

import "fmt"

/*
 * Shadowing occurs in Go when a variable declared in a nested scope has the same name as a variable declared in an outer scope.
 * In such cases, the variable in the inner scope shadows or hides the variable in the outer scope.
 */

var a = 10

func shadow() {
	age := 30

	if age > 18 {
		a := 40
		fmt.Println("Now the shadowing value of is: ", a)
	}

	fmt.Println(a)
}