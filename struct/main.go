package main

import "fmt"

type User struct {
	Name  string // Property
	Age   int
	Email string
}

func printUserDetails(user User) {
	fmt.Println("Name: ", user.Name)
	fmt.Println("Age: ", user.Age)
	fmt.Println("Email: ", user.Email)
}

func main() {
	var user_1 User

	user_1 = User{
		Name:  "Ali",
		Age:   24,
		Email: "ali@gmail.com",
	}

	printUserDetails(user_1)

	fmt.Println("-------- Second User -----------")

	user_2 := User{ // Instance or Object
		Name:  "Akkas",
		Age:   25,
		Email: "akkas@gmail.com",
	}

	printUserDetails(user_2)
}

/*
	2 phases:
		1. Compilation phase (compile time)
		2. Execution phase (run time)



	************  Compilation Phase (compile time)  ***************

		### Code Segment (readonly => (const and function)) ###
			- User = type User struct {...}
			- printUserDetails = func () {...}
			- main



	************  Execution Phase (run time)  ***************




	go build main.go -> compile it -> main
	./main
*/
