package main

import "fmt"

func main() {
	temp := 10

	if temp > 5 {
		fmt.Println("Condition Satisfied") // True
	}

	marks := 60

	if marks > 40 {
		fmt.Println("Passed") // True
	} else {
		fmt.Println("Failed")
	}

	grade := ""

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

	fmt.Println(grade) // B

	if err := process(); err != nil {
		fmt.Println(err)
	}

	// Switch Statements
	temperature := 30

	switch {
	case temperature > 30:
		fmt.Println("Hot weather")
	case temperature >= 25:
		fmt.Println("Room temperature")
	default:
		fmt.Println("Cold weather")
	}

	day := "Friday"
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Weekend coming")
	default:
		fmt.Println("Weekend!")
	}

	// fallthrouth in go
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

	// For loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	count := 5
	for count < 5 {
		count++
	}

	// Loop on array/slice
	arr := []int{1, 2, 3, 4, 5}
	for i, val := range arr {
		fmt.Printf("%d: %d\n", i, val)
	}

	// Loop on map
	m := map[string]int{"x": 10, "y": 20, "z": 30}
	for key, val := range m {
		fmt.Printf("%s: %d\n", key, val)
	}

	// Loop on string
	str := "Akkas"
	for i, val := range str {
		fmt.Printf("%d: %s\n", i, string(val))
	}

	// Labels with break and continue
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 3 {
				break outer
			}
			fmt.Println("i:", i, "j:", j)
		}
	}

}

// process is a sample function that returns nil error.
func process() error {
	return nil
}
