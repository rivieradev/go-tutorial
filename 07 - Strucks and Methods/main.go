package main

import "fmt"

// Define a struct
type Person struct {
	Name string
	Age  int
	Email string
}

// Method on struct
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old. You can reach me at %s.", p.Name, p.Age, p.Email)
}

// Method with pointer receiver (can modify struct)
func (p *Person) HaveBirthday() {
	p.Age++
}

func main() {
	// Create struct instances
	person1 := Person{
		Name: "Alice",
		Age: 25,
		Email: "alice@example.com",
	}

	person2 := Person {"Bob", 30, "bob@example.com"}

	fmt.Println(person1)
	fmt.Println(person2)

	//Call methods
	fmt.Println(person1.Greet())

	person1.HaveBirthday()
	fmt.Printf("%s is now %d years old.\n", person1.Name, person1.Age)
}