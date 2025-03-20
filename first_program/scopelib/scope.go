package scopelib

import "fmt"

/*
 * Block -> { }
 * Local Scope
 * Global Scope

 * go mod init example.com (create a new module)

 * Memory allocation start for the Global Scope at begining of the program execution
 * Then program start to execute from the main function -> main()
 * Firstly check the Local scope, if not found then check the Global scope
 */

 var p = 3
 var q = 7

func Scope() {
	x := 18
    if x >= 18 {
        y := 10

        fmt.Println("I am matured because my age is: ", x)
        fmt.Println("Now I have", y, "lands")
        fmt.Println("And", p ,"wifes")
    }
}