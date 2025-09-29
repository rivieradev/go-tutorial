package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	result1 := sum(1, 2, 3)
	result2 := sum(1, 2, 3, 4, 5)

	fmt.Println("Sum of 1, 2, 3:", result1)
	fmt.Println("Sum of 1, 2, 3, 4, 5:", result2)

	// Pass slice to variadic function
	numbers := []int{10, 20, 30}
	result3 := sum(numbers...)
	fmt.Println("Sum of slice:", result3)
}