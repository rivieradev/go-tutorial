package main

import (
	"fmt"
	"math"
)

// Define an interface

type Shape interface {
	Area() float64
	Perimeter() float64
}

// Reactangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle struct
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Function that works with any Shape
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 4}

	fmt.Println("Rectangle:")
	PrintShapeInfo(rect)

	fmt.Println("Circle:")
	PrintShapeInfo(circle)

	// Interface slice
	shapes := []Shape{rect, circle}

	fmt.Println("\nAll Shapes:")
	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		PrintShapeInfo(shape)
	}
}
