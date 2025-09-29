package main

import "fmt"

func main() {
	// Create slice
	slice1 := []int{1, 2, 3, 4, 5 }

	// Create slice with make
	slice2 := make([]int, 3, 5) // length 3, capacity 5

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice2), cap(slice2))

	// Append to slice
	slice1 = append(slice1, 6, 7)
	fmt.Println("After appending to Slice 1:", slice1)

	// Slice operations
	subSlice := slice1[1:4] // Elements from index 1 to 3
	fmt.Println("Sub-slice of Slice 1 (index 1 to 3):", subSlice)

	// Copy slices
	slice3 := make([]int, len(slice1))
	copy(slice3, slice1)
	fmt.Println("Copied Slice 3 from Slice 1:", slice3)

	// Modify original slice to show that slice3 is a separate copy
	slice1[0] = 100
	fmt.Println("After modifying Slice 1:")
	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 3 (should be unchanged):", slice3)
}