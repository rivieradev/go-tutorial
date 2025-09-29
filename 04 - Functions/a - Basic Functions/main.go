package main

import "fmt"

// Function with parameters and return value
func add (a int, b int) int {
	return a + b
}

// Multiple parameters of same type
func multiply (x, y int) int {
	return x * y
}

// Multiple return values
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Named return values
func rectange(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return // naked return
}

func main() {
	sum := add(3, 5)
	fmt.Println("Sum:", sum)

	product := multiply(4, 7)
	fmt.Println("Product:", product)

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", result)
	}

	area, perimeter := rectange(5, 3)
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", area, perimeter)
}