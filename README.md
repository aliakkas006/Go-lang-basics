# 📘Go Lang Basics

Welcome to the **Go Basics Repository** – a concise and practical guide for learning the foundational concepts of the Go programming language.

This repository includes clear examples and explanations for the following core Go concepts:

## 📚 Index

- [Variables and Data Types](#-variables-and-data-types)
- [Scope](#-scope)
- [Closure](#closure)
- [Struct](#struct)
- [Array](#array)
- [Pointer](#pointer)
- [Slice](#slice)

---

# 📦 Variables and Data Types

Go is a statically typed, compiled language with a rich set of built-in data types and flexible variable declaration syntax.

## 📘 Declaring Variables in Go

- ### Explicit Declaration with `var` keyword

```go
var name string       // Declaration without initialization (default: "")
var age int = 24      // Declaration with initialization
var height = 165.5    // Type inference (float64)
```

- ### Short Declaration

```go
name := "Ali Akkas"       // string
age := 24             // int
height := 165.5        // float64
active := true        // bool
```

- ### Multiple Variable Declaration 🧩

```go
var (
	name = "ali"
	age  = 24
)

x, y := 10, 20
```

### 🧊 Zero Values

| Type      | Zero Value   |
| --------- | ------------ |
| int       | `0`          |
| float64   | `0.0`        |
| string    | `""` (empty) |
| bool      | `false`      |
| pointer   | `nil`        |
| slice/map | `nil`        |

---

## 🧬 Data Types in Go

> Go is statically typed — each variable must have a defined type (either explicitly or inferred).

### 🔢 Basic Types

| Type                                            | Description     | Example                     |
| ----------------------------------------------- | --------------- | --------------------------- |
| `int`, `int8`, `int16`, `int32` `int64`, `uint` | Integer types   | `var age int = 30`          |
| `float32`, `float64`                            | Floating points | `var price float64 = 9.99`  |
| `string`                                        | Text (UTF-8)    | `var name string = "Go"`    |
| `bool`                                          | Boolean values  | `var valid bool = true`     |
| `complex64`, `complex128`                       | Complex numbers | `var c complex128 = 1 + 2i` |

### 🏗️ Composite Types

| Type        | Description                        | Example                                       |
| ----------- | ---------------------------------- | --------------------------------------------- |
| `array`     | Fixed-size sequence of elements    | `arr := [3]int{1, 2, 3}`                      |
| `slice`     | Dynamic-length version of array    | `s := []int{1, 2, 3}`                         |
| `map`       | Key-value store                    | `m := map[string]int{"a": 1}`                 |
| `struct`    | Collection of fields (custom type) | `type Person struct { Name string; Age int }` |
| `pointer`   | Holds memory address of a value    | `var p *int`                                  |
| `interface` | Defines behavior (methods)         | `type Shape interface { Area() float64 }`     |
| `channel`   | Communication between goroutines   | `ch := make(chan int)`                        |
| `function`  | First-class functions              | `func add(a, b int) int { return a + b }`     |

### 🧪 Type Conversion

- Go is **strictly typed** — implicit conversion is not allowed.
- Use explicit conversion with the type name.

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// String to bytes and back
str := "hello"
bytes := []byte(str)
newStr := string(bytes)
```

### 🧠 Constants

- Use `const` to declare immutable values

```go
const Pi = 3.1415
const Lang = "Go"
```

---

# 🧭 Scope

Scope defines the **region of code** where a **variable**, **function**, or **other identifier** can be **accessed**.

> It determines the visibility and lifetime of program elements, preventing naming conflicts and managing memory efficiently.

There are **four primary scopes** in Go:

1. **Block Scope**
2. **Function Scope**
3. **Package Scope**
4. **File Scope**

## 1️⃣ Block Scope

- Variables declared inside **`{}`** blocks are only visible within that block
- Includes control structures (**`if`**, **`for`**, **`switch`**), **`functions`**, and **`explicit blocks`**

```go
func main() {
    x := 10 	// Outer scope variable

    {
        y := 30 	// Inner scope variable
        fmt.Println(x, y) 	// Valid (10 30)
    }

    fmt.Println(y) 	// Compile error: undefined:y
}
```

## 2️⃣ Function Scope

- Variables declared inside functions are visible only within that function.
- Includes parameters and return values.

```go
func calculate(a, b int) (result int) {		// a, b, and result are function-scoped
    temp := a * b 	// Also function-scoped
    return temp + 10
}

func main() {
	result := calculate(10, 20)
	fmt.Println(result) // 210
    fmt.Println(temp) // Error: undefined:temp
}
```

> ⚠️ These variables are re-created every time the function is called.

## 3️⃣ Package Scope

- Variables declared **outside of functions** are accessible **anywhere in the same package**.
- Can be exported (visible to other packages) if capitalized.

```go
package main

var msg = "Package-scoped variable" // Not exported
var Msg = "Package-scoped variable" // Exported

func show() {	// Not Exported
    fmt.Println(msg) // Works
}

func Add(x, y int) int {	// Exported
	return x + y
}

func main() {
    fmt.Println(msg) // Works
}
```

> ✅ Package-scoped variables **persist for the lifetime** of the program.
> ⚠️ Can lead to **race conditions** in concurrent code.

## 4️⃣ File Scope

- Imported packages are file-scoped.

```go
package main

import (
    "fmt"  // Only visible in this file
    m "math"  // File-scoped alias
)

var _ = secretHelper() // File-scoped initialization

func main() {
    fmt.Println(m.Sqrt(4))
}
```

## 🔄 Variable Shadowing

- When an **inner scope variable** declares **same name** as **outer scope variable**.

```go
x := "Outer-scope variable"
{
    x := "Inner-scope variable" 	// Shadows outer scope variable x
    fmt.Println(x) 	// Inner-scope variable
}
fmt.Println(x) 	// Outer-scope variable
```

> ⚠️ Shadowing can cause bugs — avoid reusing variable names in nested scopes unless intentional.

## 🧪 Example: Demonstrating All Scopes

```go
package main

import (
    "fmt"  // Only visible in this file
    m "math"  // File-scoped alias
)

var global = "package-scope" 	// package scope

func main() {
	local := "function-scope"

	if true {
		block := "block-scope"

		fmt.Println(global)		// ✅ Can access global variable
		fmt.Println(local)		// ✅ Can access local variable
		fmt.Println(block)		// ✅ Can access block variable
	}

	fmt.Println(block)  // ❌ Not accessible here

	val := m.Sqrt(4) 	// File-scope
	fmt.Println(val)
}
```

## 🧱 Lexical Scope

> Lexical scope (also called static scope) is a fundamental concept in Go that determines how and where variables, functions, and other identifiers are accessible based on their physical location in the source code. Unlike dynamic scope (which determines visibility at runtime), lexical scope is determined at compile time.

```go
func outerFunc() func() int {
    count := 0 		// Lexically-scoped to outerFunc

    return func() int {
        count++ 		// The inner function "closes over" count
        return count
    }
}

func main() {
    counter := outerFunc()
    fmt.Println(counter()) 		// 1
    fmt.Println(counter()) 		// 2 (count persists)
}
```

- The inner function maintains a reference to count (not a copy).
- This is possible because Go uses **lexical scoping**.

## 📊 Lexical vs Dynamic Scope Comparison

| Feature           | Lexical Scope (Go)      | Dynamic Scope (e.g., Bash) |
| ----------------- | ----------------------- | -------------------------- |
| 🔍 Resolution     | At compile time         | At runtime                 |
| 🧭 Access Rules   | Based on code structure | Based on call stack        |
| ⚡ Performance    | Faster (resolved early) | Slower (runtime lookup)    |
| ✅ Predictability | More predictable        | Less predictable           |

---

# Closure

Closures are a powerful concept in Go that allow functions to **remember** and **access variables** from their surrounding lexical scope, even after the outer function has finished executing. They are often used for:

- **Encapsulation** (data privacy)
- **Function factories** (dynamically generating functions)
- **Callbacks** (event handlers, async operations)
- **Stateful functions** (maintaining state between calls)

### ✅ Basic Example:

Closures in Go are **anonymous functions** that capture variables from their surrounding scope. They are useful when we want a function with persistent state.

```go
package main
import "fmt"

const a = 10
var p = 100

func outer() func() {
	money := 500
	age := 24

	show := func() {
		money = money + a + p
		fmt.Println("Money: ", money)
	}

	return show
}

func call() {
	incr_1 := outer()
	incr_1()    // 610
	incr_1()    // 720

	fmt.Println("---- Second incremetnt ------")

	incr_2 := outer()
	incr_2()    // 610
	incr_2()    // 720
}

func main(){
	call()
}

func init(){
	fmt.Println("==== Bank ====")
}
```

```
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
			### Data Segment (var)
				- p


	* go build main.go -> compile it -> main
	./main (run)
*/
```

- Variables captured by a closure are lifted to the **heap** by the escape analysis of the go compiler.
- **Escape analysis** is a process performed by the Go compiler to determine whether a variable can be allocated on the stack or must be allocated on the heap

---

### 🔒 Encapsulation (private variables)

```go
func newBankAccount(initialBalance float64) (func(float64), func() float64) {
    balance := initialBalance 	// "private" variable

    deposit := func(amount float64) {
        balance += amount
    }

    getBalance := func() float64 {
        return balance
    }

    return deposit, getBalance
}

func main() {
    deposit, getBalance := newBankAccount(100.0)
    deposit(50.0)
    fmt.Println(getBalance()) // 150.0
}
```

- balance is hidden from outside access.

---

# Struct

> Structs are one of the most important features in Go for organizing and managing data.
> A struct is a **composite data type** that groups together fields (variables) under a single name.

### 🧱 Basic syntax

```go
type StructName struct {
	Field1 datatype1
	Field2 datatype2
	Field3 datatype3
}
```

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

## 🛠️ Creating Struct instances

### 1️⃣ Field Initialization

```go
p1 := Person{
    Name: "Akkas",
    Age:  24,
    City: "Dhaka",
}
```

### 2️⃣ Positional Initialization (Order matters!)

```go
p2 := Person{"Anis", 25, "Chittigong"}  // Must follow struct field order
```

### 3️⃣ Zero-value Initialization

```go
var p3 Person 	// All fields get their zero values
fmt.Println(p3) // Output: { 0 }
```

### 4️⃣ Using **new()** (Returns a Pointer)

```go
p4 := new(Person) // p4 is a *Person (pointer)
p4.Name = "Ali"
```

---

## ✅ Accessing Struct Fields

Use the **.** (dot) operator to access fields.

```go
fmt.Println(p1.Name) // "Akkas"
p1.Age = 30          // Modify a field
```

## ✅ Nested Structs

Example: Address inside Person

```go
type Address struct {
    Street  string
    Country string
}

type Person struct {
    Name    string
    Age     int
    Address Address // Nested struct
}

// Initialization
p := Person{
    Name: "Ali Akkas",
    Age:  24,
    Address: Address{
        Street:  "123 Jatrabari",
        Country: "Bangladesh",
    },
}

// Accessing nested fields
fmt.Println(p.Address.Country) // "Bangladesh"
```

## ✅ Anonymous Structs

- One-time use.
- If we need a struct for a short-lived purpose, we can define it inline.

```go
temp := struct {
    ID    int
    Value string
}{
    ID:    1,
    Value: "test_value",
}

fmt.Println(temp.Value)  // "test_value"
```

## ✅ Comparing Structs

- Structs can be compared **only if all fields are comparable**.

```go
p1 := Person{"Akkas", 24, "Bangladesh"}
p2 := Person{"Akkas", 24, "Bangladesh"}

fmt.Println(p1 == p2) // true
```

---

### 💡Final Thoughts

Structs are the backbone of data organization in Go. They provide:

- ✅ **Type safety**
- ✅ **Flexibility** (composition over inheritance)
- ✅ **Clean code** (group related data)

---

## 📦Receiver Function

A receiver function is a function that binds to a type (usually a struct) and can be called like a method.

### 🧱 Basic Syntax

```go
func (receiver ReceiverType) MethodName(args) ReturnType {
    // Function body
}
```

### 🧪 Example

```go
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
	user := User{
		Name:  "Ali",
		Age:   24,
		Email: "ali@gmail.com",
	}

	user.printDetails()

	user_2 := User{
		Name: "Ali Akkas",
		Age:  24,
	}

	user_2.call(30)
}
```

```go
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

```

### ✅ Pointer Receiver

- func (t \*Type)
- Works on the original struct (modifiable).

```go
func (usr *User) Birthday() {
    usr.Age++ // Modifies the original User data
}

user := User{"Akkas", 24}
person.Birthday()
fmt.Println(person.Age) // 25 (changed)
```

Receiver functions are Go’s way of attaching behavior to data without full-blown OOP. They enable:

> ✅ Clean, object-like APIs
> ✅ Polymorphism via interfaces
> ✅ Explicit control over mutability

---

# Array

> Arrays in Go are fixed-size, homogeneous (same type) data structures that store elements in contiguous memory.
> Unlike slices, arrays have a static length that cannot be changed after creation.

## 🔢 Basic Syntax

```go
var arrayName [length]Type
```

## 🧩 Array Declaration & Initialization

### 1️⃣ Zero-initialized Array

```go
var arr [3]int	// [0, 0, 0]
```

### 2️⃣ Pre-filled Array

```go
arr := [3]{1, 2, 3}
```

### 3️⃣ Implicit Array (**...**)

```go
arr := [...]{1, 2, 3, 4, 5}		// length = 5
```

## ✅ Working with Arrays

- ### Accessing Elements

```go
arr := [3]int{1, 2, 3}
fmt.Println(arr[1])  // 2
```

- ### Modifying Elements

```go
arr[1] = 20		// [1, 20, 3]
```

- ### Iteration

```go
// Using for-loop
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

// Using range
for index, value := range arr {
    fmt.Printf("%d: %d\n", index, value)
}
```

- ### Multidimensional Array

```go
 matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}

fmt.Println(matrix[1][2])  // 6
```

- ### Matrix Operation

```go
 func addMatrices(a, b [2][2]int) [2][2]int {
    var result [2][2]int
    for i := range a {
        for j := range a[i] {
            result[i][j] = a[i][j] + b[i][j]
        }
    }
    return result
}
```

---

# Pointer

Pointers are a fundamental concept in Go that allow us to **directly manipulate memory addresses**.
They are essential for:

- **Efficiently passing large data** (avoid copying)
- **Modifying variables in functions**
- **Working with data structures** (linked lists, trees)
- **Interfacing with system-level code**

## 🔍 Pointer Terminology

| Concept      | Description                                   |
| ------------ | --------------------------------------------- |
| `&` operator | **Address-of** (gets the address of a value)  |
| `*` operator | **Dereference** (gets the value at a pointer) |
| `*Type`      | **Pointer to a given type**                   |

## 🎯Pointer Syntax

A pointer is a variable that stores the **memory address** of another variable.

```go
var ptr *Type  // Declares a pointer to a Type
```

```go
package main
import "fmt"

func print(numbers *[5]int) { // pointer to an array of 5 integers
	fmt.Println(numbers) // &[1 2 3 4 5]
}

func main() {
	x := 10 	// Only value hold | 10
	ptr := &x 	// Address of x | 0xc00000a0e8 | 824633759720
	val := *ptr // Value at address ptr | 10
	*ptr = 20 	// Re-assign the value of x at the address of ptr | 20

	fmt.Println("Address of x = ", ptr)		// 0xc00000a0e8
	fmt.Println("Value at address ptr = ", val)		// 10
	fmt.Println("Final Value of x = ", x)	// 20

	// Array print with the help of pointers
	arr := [5]int{1, 2, 3, 4, 5}

	print(&arr) // pass by reference
}
```

## 🧩 Pointer Mechanics:

- ### Zero Value (**nil**)

```go
var p *int
fmt.Println(p == nil)  // true
```

```go
p := new(int)  // *int pointing to 0
*p = 42        // Store 42 at the address
```

- ### Linked Data Structures

```go
type Node struct {
	Val int
	Next *Node	// Pointer to the next node
}
```

# Slice

> Slices are one of the most powerful and commonly used data structures in Go 🧠  
> Learn how slices work under the hood, how to use them effectively, and what to avoid.

---

## 🧩 What is a Slice?

A slice is a **dynamic, flexible view** into an array. Unlike arrays, slices can:

- Resize as needed
- Reference only part of the underlying data
- Share data without owning it

### 🔬 Slice Internal Structure:

```go
type slice struct {
    ptr  unsafe.Pointer
    len  int
    cap  int
}
```

## 🔬 Slice Internal Fields

- **`ptr`**: Points to the start of the underlying array segment.
- **`len`**: Number of elements the slice currently holds.
- **`cap`**: Total capacity from start pointer to the end of the array.

> 💡 Slices are reference types — copying a slice doesn't copy the data, just the reference.

---

## 🛠️ Creating Slices

### 1️⃣ Array to Slice

- **`slice := arr[start(inclusive):end(exclusive)]`**
- **`len = end - start`**
- **`cap = len(arr) - start`**

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]    // [2, 3, 4], len = 3, cap = 4
```

### 2️⃣ Slice to Slice

```go
s1 := s[1:3]    // [3, 4], len = 2, cap = 3
```

- Still points to the same array memory
- Cap is relative to original array

### 3️⃣ Slice Literals

```go
s := []int{1, 2, 3}    // [1, 2, 3], len = 3, cap = 3
```

- Creates slice with a backing array

### 4️⃣ Using **make()**

- **`slice := make([]type, len)`**
- **`slice := make([]type, len, cap)`**

```go
s := make([]int, 3) // [0, 0, 0], len = 3, cap = 3

s := make([]int, 3, 5) // [0, 0, 0], len = 3, cap = 5
s[0] = 5               // [5, 0, 0], len = 3, cap = 5
s[2] = 10              // [5, 0, 10], len = 3, cap = 5
s[3] = 20               // runtime error: index out of range [3] with length 3
```

- First form: len = cap
- Second form: custom capacity

### 5️⃣ Nil Slice

```go
var s []int  // [], ptr = nil, len = 0, cap = 0
```

- No memory allocated for the nil slice
- Common in function return signatures

### 6️⃣ Empty Slice

```go
s := []int{}    // [], len = 0, cap = 0
s := make([]int, 0)     // [], len = 0, cap = 0
```

- Yes, memory allocated for the empty slice
- Used in APIs to return empty results

---

## ➕ Appending to Slices

```go
func main() {
    var x []int     // [], ptr = nil, len = 0, cap = 0

    x = append(x, 1)    // [1], len = 1, cap = 1
    x = append(x, 2)    // [1, 2], len = 2, cap = 2
    x = append(x, 3)    // [1, 2, 3], len = 3, cap = 4

    y := x  // [1, 2, 3], len = 3, cap = 4 ; new slice (y) created and point to the (x) slice

    x = append(x, 4)    // [1, 2, 3, 4], len = 4, cap = 4
    y = append(y, 5)    // [1, 2, 3, 5], len = 4, cap = 4

    x[0] = 10

    fmt.Println("x = ", x, len(x), cap(x))  // [10, 2, 3, 5], len = 4, cap = 4
    fmt.Println("y = ", y, len(y), cap(y))  // [10, 2, 3, 5], len = 4, cap = 4
}
```

- ⚠️ Append might create a new array, but if cap is sufficient, both slices share memory!
- **x** and **y** share memory until append() causes reallocation

Go handles slice growth automatically with these rules:

| Condition  | Growth Factor |
| ---------- | ------------- |
| cap ≤ 1024 | ×2            |
| cap > 1024 | ×1.25         |

### 📤 Passing Slices to Functions

```go
func changesSlice(z []int) []int {
	z[0] = 10
	z = append(z, 11) // [10, 6, 7, 11], len = 4, cap = 6
	return z
}
```

### 🧪 Memory Sharing Demonstration

```go
func main() {
    x := []int{1, 2, 3, 4, 5}   // [1, 2, 3, 4, 5], len = 5, cap = 5

    x = append(x, 6)    // [1, 2, 3, 4, 5, 6], len = 6, cap = 10
    x = append(x, 7)    // [1, 2, 3, 4, 5, 6, 7], len = 7, cap = 10

    a := x[4:]  // [5, 6, 7], len = 3, cap = 6

    y := changesSlice(a)    // [10, 6, 7, 11], len = 4, cap = 6

    fmt.Println("x = ", x, len(x), cap(x))  // [1, 2, 3, 4, 10, 6, 7], len = 7, cap = 10
    fmt.Println("y = ", y, len(y), cap(y))  // [10, 6, 7, 11], len = 4, cap = 6

    fmt.Println("x = ", x[0:8], len(x), cap(x))     // [1, 2, 3, 4, 10, 6, 7, 11], len = 7, cap = 10
}
```

- Original array: [1, 2, 3, 4, 10, 6, 7, 11, _, _]

---

## 📌 Summary of Slice Creation

| Method                  | Allocates Memory? | Description                    |
| ----------------------- | ----------------- | ------------------------------ |
| `arr[start:end]`        | No                | Slice of existing array        |
| `s[start:end]`          | No                | Slice of slice (shares memory) |
| `s []T{}` (empty slice) | Yes               | Empty but allocated            |
| `make([]T, len)`        | Yes               | Allocates with length          |
| `make([]T, len, cap)`   | Yes               | Allocates with custom capacity |
| `var s []T` (nil slice) | No                | No allocation                  |

---

## 📦 Variadic Function in Go

- A variadic function in Go is a function that accepts a **variable number of arguments** of the same type.
- use the ... (ellipsis) syntax before the type to define it.
- example: func funcName(arg ...type)

```go
func variadicFunc(numbers ...int) { // numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers, len(numbers), cap(numbers)) // [1, 2, 3, 4, 5], len = 5, cap = 5
}
```

```go
func sum(nums ...int) int { // nums := []int{1, 2, 3}
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
```

```go
func main() {
	variadicFunc(1, 2, 3, 4, 5)

	summation2 := sum(1, 2, 3)      // 6
	summation := sum(1, 2, 3, 4, 5) // 15
}
```

## 📌 Summary of Variadic Function

- When we call a variadic function, Go converts the arguments into a slice.
- The variadic parameter is implemented as a slice under the hood.
- The compiler generates code to automatically create this slice.
- Calling with no variadic args creates a nil slice: **numbers := []int**
