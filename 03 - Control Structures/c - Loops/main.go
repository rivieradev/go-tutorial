package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Count:", i)
	}

	// While style loop
	j := 0
	for j < 3 {
		fmt.Println("While:", j)
		j++
	}

	// Infinite loop (use break to exit)
	counter := 0
	for {
		if counter >= 2 {
			break
		}
		fmt.Println("Infinite:", counter)
		counter++
	}

	// Range loop (we'll see more of this with slices)
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// Range loop ignoring index
	for _, value := range numbers {
		fmt.Println("Value only:", value)
	}
}