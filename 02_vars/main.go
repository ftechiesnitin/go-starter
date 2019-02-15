package main

import "fmt"

func main() {
	// one way to declare a variable
	var hello = "hello world v"
	var ver = 2

	// shorthand way to declare a variable inside function
	// we can not declare this way outside function scope
	newHello := "hello world version 3"

	// or for multiple variables
	firstName, lastName := "James", "cameron"

	fmt.Println(hello, ver, newHello)
	fmt.Println(firstName, lastName)

	// %T will print the data type of the variable
	fmt.Printf("%T\n", ver)
}
