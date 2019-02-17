package main

import "fmt"

func main() {
	x := 12
	y := 8

	// if else
	if x > y {
		fmt.Printf("%d is greater than %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", y, x)
	}

	// else if
	color := "red"

	if color == "red" {
		fmt.Printf("Color is %s\n", color)
	} else if color == "blue" {
		fmt.Printf("Color is %s\n", color)
	} else {
		fmt.Printf("Color is %s\n", color)
	}

	// switch statement
	switch color {
	case "red":
		fmt.Printf("Color is %s\n", color)
	case "green":
		fmt.Printf("Color is %s\n", color)
	default:
		fmt.Printf("Color is %s\n", color)
	}
}
