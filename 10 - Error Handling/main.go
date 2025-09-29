package main

import (
	"errors"
	"fmt"
)

// Custom error type
type ValidationError struct {
	Field   string
	Message string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("Validation error on field '%s': %s", v.Field, v.Message)
}

// Function with custom error
func validateAge(age int) error {
	if age < 0 {
		return ValidationError{
			Field:   "age",
			Message: "age cannot be negative",
		}
	}
	if age > 150 {
		return ValidationError{
			Field:   "age",
			Message: "age seems unrealistic",
		}
	}
	return nil
}

// Function that returns an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// Handle standard errors
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Division result:", result)

	// Handle division by zero
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Handle custom errors
	ages := []int{25, -5, 200}

	for _, age := range ages {
		if err := validateAge(age); err != nil {
			fmt.Printf("Age %d: %s\n", age, err)
		} else {
			fmt.Printf("Age %d: valid\n", age)
		}
	}
}
