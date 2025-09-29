package main

import "fmt"

func main() {
	age := 18

	if age >= 18 {
		fmt.Println("You are an adult")
	} else if age >= 13 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a child")
	}

	// If with initialization
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	}
}