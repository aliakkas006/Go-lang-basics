# 🚀 Go In-depth Tour 🧭

Welcome to the **Go In-Depth Tour** — your one-stop repository for mastering the **Go programming language** from the ground up to advanced concepts. Whether you're a beginner exploring Go's syntax or an experienced developer diving into **low-level memory management** and **closures**, this repo is crafted to help you learn **practically**, **visually**, and **thoroughly**.

This repository includes clear examples and explanations for the following core Go concepts:

## 📚 Index

- [Keywords 🔑](#-keywords-in-go)
  - [Declaration Keywords](#-declaration-keywords)
  - [Control Flow Keywords](#-control-flow-keywords)
  - [Memory & Concurrency Keywords](#-memory--concurrency-keywords)
  - [Miscellaneous Keywords](#️-miscellaneous-keywords)
  - [Complete Keyword List](#complete-keyword-list)
  - [Go Built-in Functions](#️-go-built-in-functions)
- [Operators 🧮](#-operators)
  - [Arithmetic Operators](#-arithmetic-operators)
  - [Comparison Operators](#-comparison-operators)
  - [Logical Operators](#-logical-operators)
  - [Bitwise Operators](#-bitwise-operators)
  - [Assignment Operators](#-assignment-operators)
  - [Address & Pointer Operators](#-address--pointer-operators)
  - [Channel Operators](#-channel-operators)
- [Variables and Data Types 📦](#-variables-and-data-types)
  - [Declaring Variables](#-declaring-variables)
  - [Data Types](#-data-types)
  - [Type Conversion](#-type-conversion)
- [Control Statements 🔁](#-control-statements)
  - [Conditional Statements](#1️⃣-conditional-statements)
  - [Switch Statement](#2️⃣-switch-statement)
  - [Loop Control Statements](#3️⃣-loop-control-statements)
  - [Control Flow Statements](#4️⃣-control-flow-statements)
  - [`defer` Statement](#-defer-statement)
- [Functions 🔧](#-functions)
  - [Function Syntax](#-function-syntax)
  - [Function with Multiple Return Values](#-function-with-multiple-return-values)
  - [Named Return Values](#-named-return-values)
  - [Recursive Functions](#-recursive-functions)
  - [Anonymous Functions](#-anonymous-functions)
  - [IIFE](#-iife--immediately-invoked-function-expression)
  - [Closures](#-closures)
  - [`defer` with Functions](#-defer-with-functions)
  - [`init` function](#️-init-function)
  - [Variadic Function](#-variadic-function)
  - [Higher Order Function (HOF)](#-higher-order-function-hof)
  - [Receiver Function](#receiver-function)
- [Scope 🧭](#-scope)
  - [Block Scope](#1️⃣-block-scope)
  - [Function Scope](#2️⃣-function-scope)
  - [Package Scope](#3️⃣-package-scope)
  - [File Scope](#4️⃣-file-scope)
  - [Variable Shadowing](#-variable-shadowing)
  - [Lexical Scope](#-lexical-scope)
  - [Demonstrating All Scopes](#-demonstrating-all-scopes)
- [Closure 🧠](#-closure)
  - [Basic Example with Memory Management](#-basic-example)
  - [Encapsulation](#-encapsulation-private-variables)
- [Struct 🧱](#-struct)
  - [Basic Syntax](#-basic-syntax-1)
  - [Creating Struct instances](#️-creating-struct-instances)
  - [Accessing Struct Fields](#-accessing-struct-fields)
  - [Nested Structs](#-nested-structs)
  - [Anonymous Structs](#-anonymous-structs)
  - [Comparing Structs](#-comparing-structs)
- [Array 🔢](#-array)
  - [Basic Syntax](#-basic-syntax-2)
  - [Array Declaration & Initialization](#-array-declaration--initialization)
  - [Working with Arrays](#-working-with-arrays)
- [Pointer 🎯](#-pointer)
  - [Pointer Terminology](#-pointer-terminology)
  - [Pointer Syntax](#pointer-syntax)
  - [Pointer Mechanics](#-pointer-mechanics)
- [Slice 🧩](#-slice)
  - [Slice Structure](#️-slice-structure)
  - [Creating Slices](#️-creating-slices)
  - [Appending to Slices](#-appending-to-slices)
  - [Summary of Slice Creation](#-summary-of-slice-creation)
- [Map 🗺️](#️-maps)
  - [Map Syntax](#-map-syntax)
  - [Declaring Maps](#️-declaring-maps)
  - [Working with Map](#-working-with-map)
  - [Practical Example](#-practical-example-word-frequency-counter)
- [Contributions](#-contributions)

---

# Keywords in Go

> In Go, keywords are reserved words that have special meaning in the language. we can't use them as variable names or identifiers. Go has just **25 keywords** that form the building blocks of the language.

## Declaration Keywords

| Keyword   | Purpose                           |
| --------- | --------------------------------- |
| `var`     | Declares variables                |
| `const`   | Declares constants                |
| `type`    | Declares a new type               |
| `func`    | Declares a function               |
| `package` | Declares the current package name |
| `import`  | Imports packages                  |

---

## Control Flow Keywords

| Keyword       | Purpose                                  |
| ------------- | ---------------------------------------- |
| `if`          | Starts a conditional block               |
| `else`        | Specifies an alternative block           |
| `switch`      | Multiple conditional branches            |
| `case`        | A branch within `switch`                 |
| `default`     | Default branch in `switch`               |
| `for`         | Starts a loop                            |
| `range`       | Iterates over arrays, slices, maps, etc. |
| `goto`        | Jumps to a label                         |
| `fallthrough` | Falls through to next case in switch     |
| `continue`    | Skips current loop iteration             |
| `break`       | Exits loop or switch                     |
| `return`      | Exits from a function                    |
| `defer`       | Delays execution until function ends     |

---

## Memory & Concurrency Keywords

| Keyword  | Purpose                       |
| -------- | ----------------------------- |
| `chan`   | Declares a channel            |
| `go`     | Starts a goroutine            |
| `select` | Waits on multiple channel ops |

---

## Miscellaneous Keywords

| Keyword     | Purpose                    |
| ----------- | -------------------------- |
| `interface` | Declares an interface type |
| `struct`    | Declares a structure type  |
| `map`       | Declares a map type        |

---

## Complete Keyword List

| Keyword    | Keyword       | Keyword  | Keyword     | Keyword  |
| ---------- | ------------- | -------- | ----------- | -------- |
| `break`    | `default`     | `func`   | `interface` | `select` |
| `case`     | `defer`       | `go`     | `map`       | `struct` |
| `chan`     | `else`        | `goto`   | `package`   | `switch` |
| `const`    | `fallthrough` | `if`     | `range`     | `type`   |
| `continue` | `for`         | `import` | `return`    | `var`    |

#### ❗ Notes:

- Go also reserves some **predeclared identifiers** like `int`, `string`, `true`, `false`, `nil`, etc., which aren't keywords but are important.
- Go has no classes, try-catch, or while—this keeps the language simple.
- **Case-sensitive**: all keywords must be lowercase.

---

## Go Built-in Functions

- These identifiers — `append`, `cap`, `close`, `complex`, `copy`, `delete`, `imag`, `len`, `make`, `new`, `panic`, `print`, `println`, `real`, `recover` — are **predeclared built-in functions** in Go.
- They are **not keywords**, but they're treated specially by the compiler.
- These functions are part of Go's **runtime** and are always available without importing any package.

| Function  | Purpose                                                            |
| --------- | ------------------------------------------------------------------ |
| `append`  | Adds elements to the end of a slice and returns the updated slice  |
| `cap`     | Returns the capacity of a slice, array, or channel                 |
| `close`   | Closes a channel                                                   |
| `complex` | Creates a complex number from two float values                     |
| `copy`    | Copies elements from a source slice to a destination slice         |
| `delete`  | Deletes a key-value pair from a map                                |
| `imag`    | Returns the imaginary part of a complex number                     |
| `len`     | Returns the length of strings, arrays, slices, maps, or channels   |
| `make`    | Allocates and initializes slices, maps, or channels                |
| `new`     | Allocates memory for a variable and returns a pointer to it        |
| `panic`   | Stops the normal execution of a program (used for critical errors) |
| `print`   | Prints values to the standard output (less commonly used)          |
| `println` | Like `print`, but adds spaces and a newline                        |
| `real`    | Returns the real part of a complex number                          |
| `recover` | Regains control of a panicking goroutine (used with `defer`)       |

---

# Operators

> Operators are special symbols used to perform operations on variables and values. They are categorized as follows:

## Arithmetic Operators

| **Operator** | **Description**          | **Example** |
| ------------ | ------------------------ | ----------- |
| `+`          | Addition                 | `a + b`     |
| `-`          | Subtraction              | `a - b`     |
| `*`          | Multiplication           | `a * b`     |
| `/`          | Division                 | `a / b`     |
| `%`          | Modulus (Remainder)      | `a % b`     |
| `++`         | Increment (Postfix only) | `a++`       |
| `--`         | Decrement (Postfix only) | `a--`       |

### ⚠️ Caution:

- Go does **not** support prefix increment/decrement (`++a`, `--a`) — only **postfix** is allowed (`a++`, `a--`).

---

## Comparison Operators

| **Operator** | **Description**          | **Example** |
| ------------ | ------------------------ | ----------- |
| `==`         | Equal to                 | `a == b`    |
| `!=`         | Not equal to             | `a != b`    |
| `>`          | Greater than             | `a > b`     |
| `<`          | Less than                | `a < b`     |
| `>=`         | Greater than or equal to | `a >= b`    |
| `<=`         | Less than or equal to    | `a <= b`    |

### Note:

- These operators return `true` or `false` and are commonly used in conditional statements.

---

## Logical Operators

| **Operator** | **Description** | **Example** |
| ------------ | --------------- | ----------- |
| `&&`         | Logical AND     | `a && b`    |
| `\|\|`       | Logical OR      | `a \|\| b`  |
| `!`          | Logical NOT     | `!a`        |

### Notes:

- `&&` returns `true` if **both** operands are `true`.
- `||` returns `true` if **at least one** operand is `true`.
- `!` inverts the boolean value (i.e., `!true` is `false`).
- Logical expressions are **short-circuited**:
  - `a && b` stops evaluating if `a` is `false`
  - `a || b` stops evaluating if `a` is `true`

---

## Bitwise Operators

| **Operator**        | **Description**     | **Example**             |
| ------------------- | ------------------- | ----------------------- |
| `&`                 | Bitwise AND         | `a & b`                 |
| <code>&#124;</code> | Bitwise OR          | <code>a &#124; b</code> |
| `^`                 | Bitwise XOR         | `a ^ b`                 |
| `&^`                | Bit clear (AND NOT) | `a &^ b`                |
| `<<`                | Left shift          | `a << b`                |
| `>>`                | Right shift         | `a >> b`                |

---

## Assignment Operators

| **Operator**         | **Description**        | **Example**              |
| -------------------- | ---------------------- | ------------------------ |
| `=`                  | Simple assignment      | `a = b`                  |
| `+=`                 | Add and assign         | `a += b`                 |
| `-=`                 | Subtract and assign    | `a -= b`                 |
| `*=`                 | Multiply and assign    | `a *= b`                 |
| `/=`                 | Divide and assign      | `a /= b`                 |
| `%=`                 | Modulus and assign     | `a %= b`                 |
| `&=`                 | Bitwise AND and assign | `a &= b`                 |
| <code>&#124;=</code> | Bitwise OR and assign  | <code>a &#124;= b</code> |
| `^=`                 | Bitwise XOR and assign | `a ^= b`                 |
| `<<=`                | Left shift and assign  | `a <<= b`                |
| `>>=`                | Right shift and assign | `a >>= b`                |

---

## Address & Pointer Operators

| **Operator** | **Description**     | **Example** |
| ------------ | ------------------- | ----------- |
| `&`          | Address of          | `&a`        |
| `*`          | Pointer dereference | `*ptr`      |

### Notes:

- `&a` returns the memory address of variable `a`.
- `*ptr` accesses the value stored at the memory address held by `ptr`.

---

## Channel Operators

| **Operator** | **Description**             | **Example**           |
| ------------ | --------------------------- | --------------------- |
| `<-`         | Send/Receive from a channel | `ch <- x`, `x = <-ch` |

### Notes:

- `ch <- x` sends the value `x` into the channel `ch`.
- `x = <-ch` receives a value from the channel `ch` and stores it in `x`.

---

# Variables and Data Types

Go is a statically typed, compiled language with a rich set of built-in data types and flexible variable declaration syntax.

## Declaring Variables

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

### Zero Values

| Type      | Zero Value   |
| --------- | ------------ |
| int       | `0`          |
| float64   | `0.0`        |
| string    | `""` (empty) |
| bool      | `false`      |
| pointer   | `nil`        |
| slice/map | `nil`        |

---

## Data Types

> Go is statically typed — each variable must have a defined type (either explicitly or inferred).

### Basic Types

| Type                                            | Description     | Example                     |
| ----------------------------------------------- | --------------- | --------------------------- |
| `int`, `int8`, `int16`, `int32` `int64`, `uint` | Integer types   | `var age int = 30`          |
| `float32`, `float64`                            | Floating points | `var price float64 = 9.99`  |
| `string`                                        | Text (UTF-8)    | `var name string = "Go"`    |
| `bool`                                          | Boolean values  | `var valid bool = true`     |
| `complex64`, `complex128`                       | Complex numbers | `var c complex128 = 1 + 2i` |

### Composite Types

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

### Constants

- Use `const` to declare immutable values

```go
const Pi = 3.1415
const Lang = "Go"
```

## Type Conversion

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

---

# Control Statements

> Control statements are fundamental building blocks that determine the flow of program execution.
> Control statements in Go allow us to manage the flow of execution within loops and functions.
> They provide ways to **skip iterations**, **exit loops early**, **jump to labeled code**, and **return from functions**.

## Conditional Statements

- ### `if` Statement

```go
if temp > 5 {

	fmt.Println("Condition Satisfied!")
}
```

- ### `if-else` Statement

```go
if marks > 40 {
	fmt.Println("Passed!")
} else {
	fmt.Println("Failed!")
}
```

- ### `if else-if` Ladder

```go
if marks >= 80 {
	grade = "A+"
} else if marks >= 70 {
	grade = "A"
} else if marks >= 60 {
	grade = "B"
} else if marks >= 40 {
	grade = "C"
} else {
	grade = "F"
}
```

- ### Short Statement

Go allows a short statement to execute before the condition.

```go
if err := process(); err != nil {
	fmt.Println(err)
}
```

## Switch Statement

- ### Basic Switch

```go
day := "Friday"
switch day {
case "Monday":
	fmt.Println("Start of the week")
case "Friday":
	fmt.Println("Weekend coming")
default:
	fmt.Println("Weekend!")
}
```

- ### Switch with `no expression`

```go
temperature := 30
switch {
case temperature > 30:
	fmt.Println("Hot weather")
case temperature >= 25:
	fmt.Println("Room temperature")
default:
	fmt.Println("Cold weather")
	}
```

- ### `Fallthrough` within switch

> In Go, the fallthrough keyword is used within a switch statement to force execution to continue to the next case, even if the next case expression does not match.

```go
switch num := 1; num {
case 1:
	fmt.Println("Case 1")
	fallthrough
case 2:
	fmt.Println("Case 2")
	fallthrough
case 3:
	fmt.Println("Case 3")
}
```

```bash
Output:
Case 1
Case 2
Case 3
```

> ⚠️ fallthrough ignores the next case's condition and just executes its body.

## Loop Control Statements

- ### Traditional For Loop

Syntax:

```go
for init; condition; inc/dec {
    ---
}
```

Example:

```go
for i := 0; i < 5; i++ {
	fmt.Println(i)
}
```

- ### While-style Loop

Syntax:

```go
for condition {
	---
}
```

Example:

```go
count := 5
for count < 5 {
	count++
}
```

- ### Infinite Loop

Syntax:

```go
for {
	---
	break	// needed to exit
}
```

Example:

```go
for {
    data := readData()
    if data == nil {
        break
    }
    process(data)
}
```

- ### Range-based Loop

Syntax:

```go
for index, value := range arr {
    ---
}
```

Example:

```go
// Loop on array/slice
arr := []int{1, 2, 3, 4, 5}

for i, val := range arr {
	fmt.Printf("%d: %d\n", i, val)
}
```

```go
// Loop on map
m := map[string]int{"x": 10, "y": 20, "z": 30}

for key, val := range m {
	fmt.Printf("%s: %d\n", key, val)
}
```

```go
// Loop on string
str := "Akkas"

for i, val := range str {
	fmt.Printf("%d: %s\n", i, string(val))
}
```

```go
// Channel
for item := range channel {
    process(item)
}
```

## Control Flow Statements

- ### `break`

```go
for i := 0; ; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
```

- ### `continue`

```go
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue
    }
    fmt.Println(i)
}
```

- ### Labels with break and continue

```go
outer:
for i := 1; i <= 3; i++ {
	for j := 1; j <= 3; j++ {
		if i*j > 4 {
			break outer
		}
		fmt.Println("i:", i, "j:", j)
	}
}
```

- ### `goto`: Jump to Labeled Code

```go
func main() {
	count := 0
start:
	if count < 3 {
		fmt.Println("Count:", count)
		count++
		goto start
	}
}
```

- ### return: Exit from Function

```go
func greet(name string) {
	if name == "" {
		return
	}
}
```

### Summary Table

| Statement  | Description                                     |
| ---------- | ----------------------------------------------- |
| `break`    | Exits current loop or switch                    |
| `continue` | Skips current iteration of loop                 |
| `goto`     | Jumps to a labeled line of code                 |
| `return`   | Exits the current function                      |
| Label      | Named target for `goto`, `break`, or `continue` |

## ⚡ `defer` Statement

```go
func readFile() {
    f, err := os.Open("file.txt")

    if err != nil {
        return
    }

    defer f.Close() // Executes when function exits

    // Process file
}
```

- Defers execute in LIFO order
- Arguments are evaluated immediately
- Useful for resource cleanup

---

# 🔧 Functions

- A function is a reusable block of code that performs a specific task.
- A parameter is a variable named in the function definition.
- An argument is the actual value that is passed to the function when it is called.

## Function Syntax

```go
func functionName(param1 type1, param2 type2) returnType {
	// --- function body
	return value
}
```

Example:

```go
func add(a int, b int) int {
	return a + b
}
```

## Function with Multiple Return Values

```go
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}
```

## Named Return Values

```go
func getStats(nums []int) (sum int, count int) {
	for _, n := range nums {
		sum += n
	}
	count = len(nums)

	return 		// Automatically returns named values - (sum, count)
}
```

## Recursive Functions

- A function that calls itself.

```go
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
```

## Anonymous Functions

- An anonymous function is a function that doesn’t have a name.
- It is useful when you want to create an inline function.
- We Can assign an anonymous function to a variable.

```go
add := func(x, y int) int {
	return x + y
}

add(2, 3) 	// 5
```

## IIFE : Immediately Invoked Function Expression

```go
func(x, y int) {
	sum := x + y
	fmt.Println(sum)
}()
```

## Closures

- Closure allows functions to **remember** and **access variables** from their surrounding **lexical scope**, even after the outer function has finished executing.

```go
func counter() func() int {
	i := 0

	op := func() int {
		i++
		return i
	}

	return op
}

next := counter()

next()	// 1
next()  // 2
```

## `defer` with Functions

- Use `defer` to delay execution until the surrounding function returns.

```go
func process() {
	defer fmt.Println("Finished!")

	fmt.Println("Processing...")
}
```

Output:

```bash
Processing...
Finished!
```

## `init` function

- We can't call this function, Computer calls this function autometically.
- It will called at the beginning of the program execution (even before main function's called).

Example:

```go
package main
import "fmt"

func init() {
	fmt.Println("init function executed")
}

func main() {
	fmt.Println("main function executed")
}
```

🟢 Output:

```bash
init function executed
main function executed
```

### Key Characteristics of `init()`

| Property             | Description                                        |
| -------------------- | -------------------------------------------------- |
| Signature            | `func init()` — no parameters, no return value     |
| Automatic Invocation | Called before `main()` and after global variables  |
| Multiple `init`s     | A package can have **multiple** `init()` functions |
| File Order           | Run in the order files are compiled                |
| Package Order        | Dependencies' `init()` run **before** yours        |

---

## Variadic Function

- A variadic function is a function that accepts a **variable number of arguments** of the same data type.
- use the ... (ellipsis) syntax before the type to define it.

Basic Syntax:

```go
func funcName(params ...type) {
	// --- function body
	return
}
```
Example:

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

	summation := sum(1, 2, 3)      // 6
	summation2 := sum(1, 2, 3, 4, 5) // 15
}
```

### Summary of Variadic Function

- When we call a variadic function, Go converts the **arguments into a slice**.
- The variadic parameter is **implemented as a slice under the hood**.
- The compiler generates code to automatically create this slice.
- Calling with no variadic args creates a nil slice: **numbers := []int**

---

## Higher Order Function (HOF)

- A function that takes another **function as a parameter** or **returns a function** as a result or does **both** is called a higher-order function.

```go
add := func(a, b int) {
	sum := a + b
	println(sum)
}
```

```go
func processOperation(a, b int, cb func(x, y int)) func(x, y int) {
	// Execute op func
	cb(a, b)

	return func div(x, y int) {
	res := x / y
	fmt.Println(res)
	}
}
```

```go
func main() {
	result := processOperation(10, 20, add)		// Executes: add(10, 20) → prints 30
	result(20, 10)	// Executes: div(20, 10) → prints 2
}
```

### Why Use Higher-Order Functions?

| Benefit            | Description                                        |
| ------------------ | -------------------------------------------------- |
| Reusability     | Abstract repeated patterns like filtering, mapping |
| Customizability | Inject behavior as parameters                      |
| Cleaner Code    | Reduce boilerplate with functional patterns        |
| Composition     | Build complex logic from small reusable functions  |

---

## Receiver Function

- A receiver function is a function that binds to a type (usually a `struct`) and can be called like a method.
- `ReceiverType` is usually a `struct`.
- The receiver can be `value` or `pointer`.

### Basic Syntax

```go
func (receiver ReceiverType) MethodName(args) ReturnType {
    // --- Function body
}
```

### Example

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

### 1. Compilation Phase (Compile Time) 🛠️

#### Code segment (`Read-only`, contains `function definitions` and `types` )

| 📍 Address | 📜 Content                                                                                                                                                            |
| ---------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `0x0000`   | `type User struct {`<br>    `Name string`<br>    `Age int`<br>    `Email string`<br>`}`                                                                               |
| `0x0100`   | `func (user User) printDetails() {`<br>    `fmt.Println("Name:", user.Name)`<br>    `fmt.Println("Age:", user.Age)`<br>    `fmt.Println("Email:", user.Email)`<br>`}` |
| `0x0200`   | `func (user User) call(x int) {`<br>    `fmt.Println("Name:", user.Name)`<br>    `fmt.Println("Age:", user.Age)`<br>    `fmt.Println("X:", x)`<br>`}`                 |
| `0x0300`   | `func main() { ... }`                                                                                                                                                 |

### 2. Execution Phase (Run Time) 🚀

#### **Stack Frame** (grows downward, contains local variables and args)

| 📍 Address | 📦 Content                                                                                                 |
| ---------- | ---------------------------------------------------------------------------------------------------------- |
| `0xFF00`   | `user (User struct)`<br>• Name: `"Ali"` (`0xA100`)<br>• Age: `24`<br>• Email: `"ali@gmail.com"` (`0xA200`) |
| `0xFE00`   | `user_2 (User struct)`<br>• Name: `"Ali Akkas"` (`0xA300`)<br>• Age: `24`<br>• Email: `""` (nil)           |

#### Heap Memory (Dynamic Allocation)

- 🔗 Strings in Go are reference types, stored dynamically on the heap and referenced via pointers.

| 📍 Address | 🧵 Content        |
| ---------- | ----------------- |
| `0xA100`   | `"Ali"`           |
| `0xA200`   | `"ali@gmail.com"` |
| `0xA300`   | `"Ali Akkas"`     |

#### Function Call Operations

| 🧪 Operation          | 🔍 Details                                                                                                                       |
| --------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `user.printDetails()` | • Copies `user` struct from `0xFF00` into a new stack frame.<br>• Accesses `Name` and `Email` via heap at `0xA100` and `0xA200`. |
| `user_2.call(30)`     | • Copies `user_2` from `0xFE00` to the call frame.<br>• Argument `x=30` pushed to stack.<br>• Name resolved from `0xA300`.       |

### Pointer Receiver

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

> Clean, object-like APIs
> Polymorphism via interfaces
> Explicit control over mutability

### 🚫 Receiver Functions Only for Types

We cannot attach a method to a built-in type like `int`. Only **custom types (structs, named types)** can have methods.

---

# Scope

Scope defines the **region of code** where a **variable**, **function**, or **other identifier** can be **accessed**.

> It determines the visibility and lifetime of program elements, preventing naming conflicts and managing memory efficiently.

There are **four primary scopes** in Go:

1. **Block Scope**
2. **Function Scope**
3. **Package Scope**
4. **File Scope**

## Block Scope

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

---

## Function Scope

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

---

## Package Scope

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

---

## File Scope

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

---

## Variable Shadowing

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

---

## Lexical Scope

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

### Lexical vs Dynamic Scope Comparison

| Feature           | Lexical Scope (Go)      | Dynamic Scope (e.g., Bash) |
| ----------------- | ----------------------- | -------------------------- |
| Resolution     | At compile time         | At runtime                 |
| Access Rules   | Based on code structure | Based on call stack        |
| Performance    | Faster (resolved early) | Slower (runtime lookup)    |
| Predictability | More predictable        | Less predictable           |

---

## Demonstrating All Scopes

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

		fmt.Println(global)		// Can access global variable
		fmt.Println(local)		// Can access local variable
		fmt.Println(block)		// Can access block variable
	}

	fmt.Println(block)  // Not accessible here

	val := m.Sqrt(4) 	// File-scope
	fmt.Println(val)
}
```

---

# Closure

Closures are a powerful concept in Go that allow functions to **remember** and **access variables** from their surrounding lexical scope, even after the outer function has finished executing. They are often used for:

- **Encapsulation** (data privacy)
- **Function factories** (dynamically generating functions)
- **Callbacks** (event handlers, async operations)
- **Stateful functions** (maintaining state between calls)

## Basic Example:

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

### 1. Compilation Phase (Compile Time) 🛠️

#### Code segment (`Read-only`, contains `function definitions` and `types` )

| 📜 Address | 📜 Content                                      |
| ---------- | ----------------------------------------------- |
| `0x0000`   | `const a = 10`                                  |
| `0x0100`   | `func outer()` with local variables and closure |
|            | - `money := 500`                                |
|            | - `age := 24`                                   |
|            | - `show := func() { ... }`                      |
|            | - `return show`                                 |
| `0x0200`   | Anonymous closure inside `outer()`              |
|            | - `money = money + a + p`                       |
|            | - `fmt.Println("Money:", money)`                |
| `0x0300`   | `func call() { ... }`                           |
| `0x0400`   | `func main() { call() }`                        |
| `0x0500`   | `func init()`                                   |
|            | - `fmt.Println("==== Bank ====")`               |

- Variables captured by a closure are lifted to the **heap** by the escape analysis of the go compiler.
- **Escape analysis** is a process performed by the Go compiler to determine whether a variable can be allocated on the stack or must be allocated on the heap

### 2. Execution Phase (Run Time) 

- The `init()` function runs **before** `main()`
- Prints: ==== Bank ====

<!-- ####  **Stack Frame** (grows downward, contains local variables and args) -->

- Global variable `p` is initialized in the **Data Segment**.

#### Data Segment

| Address  | Content   |
| -------- | --------- |
| `0xD000` | `p = 100` |

#### First `call()` Execution

#### ➤ First `outer()` Call (creates `incr_1`)

- A new **stack frame** is created.
- Local variables are placed on the **stack**.

#### Stack (outer call 1)

| Address  | Content                                       |
| -------- | --------------------------------------------- |
| `0xF100` | `money = 500`                                 |
| `0xF108` | `age = 24`                                    |
| `0xF110` | `show (closure)` → points to code at `0x0200` |

#### Heap Allocation (Closure Capture)

Closure captures `money` and **escapes** the stack, hence moved to the **heap**.

| Address  | Content                                        |
| -------- | ---------------------------------------------- |
| `0xH100` | Closure environment (refers to `money@0xF100`) |

#### First `incr_1()` Execution

- Reads `money` from closure (heap): `500`
- Reads `a` (constant) from code segment: `10`
- Reads `p` (global var) from data segment: `100`

**New Value**: `500 + 10 + 100 = 610`

 Heap after update: 0xH100.money = 610

#### Second `incr_1()` Execution

**New Value**: `610 + 10 + 100 = 720`

 Heap: 0xH100.money = 720

---

## Encapsulation (private variables)

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

- `balance` moves from stack to heap because it's referenced after `newBankAccount` returns.
- Both closures share the same `balance` reference.
- The `balance` variable is effectively private, only accessible through the returned functions.
- `balance` is hidden from outside access.

---

# Struct

> Structs are one of the most important features in Go for organizing and managing data.
> A struct is a **composite data type** that groups together fields (variables) under a single name.

## Basic Syntax

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

## Creating Struct instances

### Field Initialization

```go
p1 := Person{
    Name: "Akkas",
    Age:  24,
    City: "Dhaka",
}
```

### Positional Initialization (Order matters!)

```go
p2 := Person{"Anis", 25, "Chittigong"}  // Must follow struct field order
```

### Zero-value Initialization

```go
var p3 Person 	// All fields get their zero values
fmt.Println(p3) // Output: { 0 }
```

### Using **new()** (Returns a Pointer)

```go
p4 := new(Person) // p4 is a *Person (pointer)
p4.Name = "Ali"
```

---

## Accessing Struct Fields

Use the **.** (dot) operator to access fields.

```go
fmt.Println(p1.Name) // "Akkas"
p1.Age = 30          // Modify a field
```

## Nested Structs

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

## Anonymous Structs

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

## Comparing Structs

- Structs can be compared **only if all fields are comparable**.

```go
p1 := Person{"Akkas", 24, "Bangladesh"}
p2 := Person{"Akkas", 24, "Bangladesh"}

fmt.Println(p1 == p2) // true
```

### 💡Final Thoughts

Structs are the backbone of data organization in Go. They provide:

- **Type safety**
- **Flexibility** (composition over inheritance)
- **Clean code** (group related data)

---

# Array

> Arrays in Go are fixed-size, homogeneous (same type) data structures that store elements in contiguous memory.
> Unlike slices, arrays have a static length that cannot be changed after creation.

## Basic Syntax

```go
var arrayName [length]Type
```

## Array Declaration & Initialization

### Zero-initialized Array

```go
var arr [3]int	// [0, 0, 0]
```

### Pre-filled Array

```go
arr := [3]{1, 2, 3}
```

### Implicit Array (**...**)

```go
arr := [...]{1, 2, 3, 4, 5}		// length = 5
```

## Working with Arrays

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

## Pointer Terminology

| Concept      | Description                                   |
| ------------ | --------------------------------------------- |
| `&` operator | **Address-of** (gets the address of a value)  |
| `*` operator | **Dereference** (gets the value at a pointer) |
| `*Type`      | **Pointer to a given type**                   |

## Pointer Syntax

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

## Pointer Mechanics:

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

## Slice Structure

A slice is a **dynamic, flexible view** into an array. Unlike arrays, slices can:

- Resize as needed
- Reference only part of the underlying data
- Share data without owning it

### Slice Internal Structure:

```go
type slice struct {
    ptr  unsafe.Pointer
    len  int
    cap  int
}
```

### 🔬 Slice Internal Fields

- **`ptr`**: Points to the start of the underlying array segment.
- **`len`**: Number of elements the slice currently holds.
- **`cap`**: Total capacity from start pointer to the end of the array.

> 💡 Slices are reference types — copying a slice doesn't copy the data, just the reference.

---

## Creating Slices

### Array to Slice

- **`slice := arr[start(inclusive):end(exclusive)]`**
- **`len = end - start`**
- **`cap = len(arr) - start`**

```go
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:4]    // [2, 3, 4], len = 3, cap = 4
```

### Slice to Slice

```go
s1 := s[1:3]    // [3, 4], len = 2, cap = 3
```

- Still points to the same array memory
- Cap is relative to original array

### Slice Literals

```go
s := []int{1, 2, 3}    // [1, 2, 3], len = 3, cap = 3
```

- Creates slice with a backing array

### Using **make()**

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

###  Nil Slice

```go
var s []int  // [], ptr = nil, len = 0, cap = 0
```

- No memory allocated for the nil slice
- Common in function return signatures

### Empty Slice

```go
s := []int{}    // [], len = 0, cap = 0
s := make([]int, 0)     // [], len = 0, cap = 0
```

- Yes, memory allocated for the empty slice
- Used in APIs to return empty results

---

## Appending to Slices

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

### Passing Slices to Functions

```go
func changesSlice(z []int) []int {
	z[0] = 10
	z = append(z, 11) // [10, 6, 7, 11], len = 4, cap = 6
	return z
}
```

### Memory Sharing Demonstration

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

## Summary of Slice Creation

| Method                  | Allocates Memory? | Description                    |
| ----------------------- | ----------------- | ------------------------------ |
| `arr[start:end]`        | No                | Slice of existing array        |
| `s[start:end]`          | No                | Slice of slice (shares memory) |
| `s []T{}` (empty slice) | Yes               | Empty but allocated            |
| `make([]T, len)`        | Yes               | Allocates with length          |
| `make([]T, len, cap)`   | Yes               | Allocates with custom capacity |
| `var s []T` (nil slice) | No                | No allocation                  |

---

# Maps

A **map** is an **unordered collection** of **key-value pairs** where:

- Each key is unique
- Values are accessed via keys
- Keys must be comparable (`==` operator)

## Map Syntax

```go
map[KeyType]ValueType
```

## Declaring Maps

- ### Using `make()` function

```go
data := make(map[string]int)
```

- ### Using Map Literals

```go
data := map[string]int{
	"ali": 24,
	"anis": 25,
	"rakib": 27,
}
```

- ### Empty Map

```go
m := make(map[string]int)		// Initialized but empty
fmt.Println(m == nil)  			// false
```

- ### Nil Map

```go
var countries map[string]string
fmt.Println(countries == nil)    	// true
```

## Working with Map

- ### Adding/Updating:

```go
data["arif"] = 25
data["ali"] = 30
```

- ### Accessing

```go
myAge := data["ali"]	// 30
```

- ### Deleting

```go
delete(data, "ali")
```

- ### Check Length

```go
totalData := len(data)
```

- ### Check If Key exists

```go
age, exists := data["arif"]
if exists {
	fmt.Println("Arif's age is: ", age)
} else {
	fmt.Println("Arif's data is not found!")
}
```

- ### Shorter version

```go
if age, exists := data["arif"]; exists {
	fmt.Println(age)
}
```

- ### Reference Type Behavior

```go
original := map[string]int{"x": 1}
ref := original
ref["x"] = 2

fmt.Println(original["x"]) // 2 (changed)
```

- ### Iteration over Map

```go
for key, value := range data {
	fmt.Println("%s: %s\n", key, value)
}

// Values only
for _, value := range colors {
    fmt.Println(value)
}
```

- ### Nested Maps

```go
students := map[string]map[string]int{
		"Akkas": {
			"Math":    90,
			"Physics": 85,
		},
		"Zidan": {
			"Math":      90,
			"Chemistry": 87,
		},
	}

fmt.Println(students["Akkas"]["Math"])  	// 90
```

## Practical Example (Word Frequency Counter)

```go
func wordCount(str string) map[string]int {
	words := strings.Fields(str)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	return counts
}
```

## Map Properties

- Keys must be of a type that supports `==` comparison (e.g., string, int).
- Maps are **reference types** (passed by reference)
- Not safe for concurrent use (use `sync.Map` for concurrency)

---

# Contributions

Contributions are welcome! Whether you're fixing bugs, improving documentation, or adding new Go examples — your help is appreciated.

### How to Contribute

1. 🍴 Fork the repository
2. 📥 Clone your forked repo
3. 🛠 Create a new branch (`git checkout -b feature-name`)
4. 💻 Make your changes
5. ✅ Commit your changes (`git commit -m "Add new Go example"`)
6. 🔄 Push to the branch (`git push origin feature-name`)
7. 📝 Create a Pull Request

Make sure your code follows Go best practices and is properly documented. All PRs are reviewed before merging.

---

⭐ **Star** this repo if you found it helpful — it motivates future improvements!
