package main

import "fmt"

func main() {
	// Arrays
	var fruitArray [2]string

	// Assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Pineapple"

	// Declare and assign
	cityArray := [2]string{"Paris", "Amsterdam"}

	// Slices
	evenNumberSlice := []int{2, 4, 8, 10, 12, 14}

	fmt.Println("Arrays: ")
	fmt.Println(fruitArray)
	fmt.Println(cityArray)

	fmt.Println("Slices: ")
	fmt.Println(evenNumberSlice)
	fmt.Println(len(evenNumberSlice))
	fmt.Println(evenNumberSlice[1:3])

}
