package main

import "fmt"

func main() {
	// Integers
	var a int = 42
	var b int8 = 27
	var c uint = 42

	// Floating-point numbers
	var pi float64 = 3.14159
	var e float32 = 2.71828

	// Boolean
	var isActive bool = true

	// String
	var message string = "Hello, Go!"

	// Character (rune)
	var char rune = 'A'

	fmt.Printf("Int: %d, Float: %.2f, Bool: %t, String: %s, Char: %c\n", a, pi, isActive, message, char)
	fmt.Printf("Int8: %d, Uint: %d, Float32: %.2f\n", b, c, e)
}