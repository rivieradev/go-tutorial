package main

import "fmt"

func main() {

	//Create map
	ages := make(map[string]int)
	ages["John"] = 25
	ages["Jane"] = 30
	ages["Bob"] = 35

	// Map literal
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	fmt.Println("Ages:", ages)
	fmt.Println("Coloers:", colors)

	// Accessing values
	johnAge := ages["John"]
	fmt.Println("John's age:", johnAge)

	// Check if key exists
	age, exists := ages["Alice"]
	if exists {
		fmt.Println("Alice's age:", age)
	} else {
		fmt.Println("Alice's age not found")
	}

	// Deleting a key
	delete(ages, "Bob")
	fmt.Println("Ages after deleting Bob:", ages)

	// Iterating over a map
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}