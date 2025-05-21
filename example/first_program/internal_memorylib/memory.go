package internal_memorylib

import "fmt"

const a = 10 // Constant

var p = 100

func Call(){
	add := func(x, y int) {
		z := x + y
		fmt.Println(z)
	}

	add(5, 6)

	add(p, a)
}

/* 
	2 phases:
		1. Compilation phase
		2. Execution phase


	go run main.go => compile it => main.exe => ./main
*/

/* 
	******* Compile phase **********

	# Code Segment (readonly)
	 - a
	 - call = func() {....}
	 - add = func() {....}
	 - main = func() {....}
	 - init = func() {....}

*/