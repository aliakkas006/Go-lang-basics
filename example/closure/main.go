package main

import "fmt"

const a = 10

var p = 100

func outer() func() {
	money := 500
	age := 24

	fmt.Println("Age: ", age)

	show := func() {
		money = money + a + p
		fmt.Println("Money: ", money)
	}

	return show
}

func call() {
	incr_1 := outer()
	incr_1() // 610
	incr_1() // 720

	fmt.Println("---- Second incremetnt ------")

	incr_2 := outer()
	incr_2()	// 610
	incr_2()	// 720
}

func main(){
	call()
}

func init(){
	fmt.Println("==== Bank ====")
}

/* 
	2 phases:
		1. Compilation phase (compile time)
		2. Execution phase (run time)



	************  Compilation Phase (compile time)  ***************

		### Code Segment (readonly => (const and function)) ###

			- a
			- outer = func() {...}
			- outerAnonymous_1 = func() {...}
			- call = func() {...}
			- main = func() {...}
			- init = func() {...}

	

	************  Execution Phase (run time)  ***************
			- p




	go build main.go -> compile it -> main 
	./main
*/