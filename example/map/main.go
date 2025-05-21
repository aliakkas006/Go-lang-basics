package main

import (
	"fmt"
	"strings"
)

func wordCount(str string) map[string]int {
	words := strings.Fields(str)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	return counts
}

func main() {
	// Map literals
	m := map[string]int{
		"ali":  24,
		"anis": 25,
	}

	fmt.Println(m)

	// Nested Map
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

	fmt.Println(students["Akkas"]["Math"]) // 90

	// Maps with struct
	type Person struct {
		City string
		Age  int
	}

	people := map[string]Person{
		"ali":  {"Dhaka", 24},
		"anis": {"Chandpur", 25},
	}

	fmt.Println(people["ali"].Age) // 24

	// Word Frequency Counter
	wordCount("akkas")
}
