package main

import "fmt"

func main() {
	// Array declaration
	var arr1 [5]int
	arr1[0] = 10
	arr1[1] = 20

	// Array initialization
	arr2 := [5]int{1, 2, 3, 4, 5}
	arr3 := [...]int{10, 20, 30} // Compiler infers size

	fmt.Println("Array 1:", arr1)
	fmt.Println("Array 2:", arr2)
	fmt.Println("Array 3:", arr3)
	fmt.Println("Length of Array 2:", len(arr2))
}