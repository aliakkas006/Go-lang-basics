package main

import "fmt"

type User struct {
	Name  string // Property
	Age   int
	Email string
}

// Receiver Function (receive only User type's variable)
func (user User) printDetails() {
	fmt.Println("Name: ", user.Name)
	fmt.Println("Age: ", user.Age)
	fmt.Println("Email: ", user.Email)
}

func (user User) call(x int) {
	fmt.Println("Name: ", user.Name)
	fmt.Println("Age: ", user.Age)

	fmt.Println("X: ", x)
}

func main() {
	var user_1 User

	user_1 = User{
		Name:  "Ali",
		Age:   24,
		Email: "ali@gmail.com",
	}

	user_1.printDetails()

	fmt.Println("-------- Second User -----------")

	user_2 := User{ // Instance or Object
		Name:  "Akkas",
		Age:   25,
		Email: "akkas@gmail.com",
	}

	user_2.printDetails()

	fmt.Println("-------- Third User -----------")

	user_3 := User{
		Name: "Ali Akkas",
		Age:  24,
	}

	user_3.call(30)
}

/*
	2 phases:
		1. Compilation phase (compile time)
		2. Execution phase (run time)



	************  Compilation Phase (compile time)  ***************

		### Code Segment (readonly => (const and function)) ###
			- User = type User struct {...}
			- printDetails = func () {...}  // Bind with User type variable
            - call = func (int) {...}   // // Bind with User type variable
			- main



	************  Execution Phase (run time)  ***************




	go build main.go -> compile it -> main
	./main
*/
