package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Email string
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old. You can reach me at %s.", p.Name, p.Age, p.Email)
}

type Address struct {
	Street  string
	City    string
	Country string
}

type Employee struct {
	Person // Embedded struct
	Address // Embedded struct
	JobTitle string
	Salary   float64
}

func main() {
	emp := Employee{
		Person: Person{
			Name:  "John Doe",
			Age:   28,
			Email: "john@doe.com",
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			Country: "USA",
		},
		JobTitle: "Software Engineer",
		Salary:   75000,
	}

	// Access embedded fields
	fmt.Println("Name:", emp.Name) // From Person
	fmt.Println("Age:", emp.City)  // From Address
	fmt.Println("Job:", emp.JobTitle)

	// Call embedded methods
	fmt.Println(emp.Greet()) // From Person
}