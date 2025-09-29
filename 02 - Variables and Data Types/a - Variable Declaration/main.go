package main

import "fmt"

func main() {
	// Method 1: Using var keyword
	var name string = "John"
	var age int = 25

	// Method 2: Type inference
	var city = "New York"

	// Method 3: Short declaration (inside functions only)
	country := "USA"

	fmt.Println(name, age, city, country)
}