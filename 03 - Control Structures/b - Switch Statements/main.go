package main

import "fmt"

func main() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the work week")
	case "Friday":
		fmt.Println("TGIF!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Midweek")
	}

	// Switch whitout expression (Like if-else chain)

	score := 85
	switch {
	case score >= 90:
		fmt.Println("Excellent")
	case score >= 80:
		fmt.Println("Good")
	default:
		fmt.Println("Needs improvement")
	}
}