# 🚀 Go Lang Basics

Welcome to the **Go Basics Repository** – a concise and practical guide for learning the foundational concepts of the Go programming language.

This repository includes clear examples and explanations for the following core Go concepts:

## 📚 Index

- [Closure](#closure)
- [Struct](#struct)
- [Array](#array)
- [Pointer](#pointer)
- [Slice](#slice)

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

### 🔒  Encapsulation (private variables)

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
