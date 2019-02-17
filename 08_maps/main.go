package main

import "fmt"

func main() {
	// defining map : Map is a key value pair
	// using make method to create map
	// first argurment of map is KEY datatype
	// next one VALUE datatype
	emails := make(map[string]string)

	// assign key value
	emails["bob"] = "bob@gmail.com"
	emails["alice"] = "alice@gmail.com"
	emails["mike"] = "mike@gmail.com"
	emails["jordon"] = "jordon@gmail.com"

	fmt.Println("Emails Map: ", emails)
	fmt.Println("Accessing a Value from map: ", emails["bob"])
	fmt.Println("Length of map: ", len(emails))

	// deleting a key
	delete(emails, "mike")
	fmt.Println("Emails Map after deleting: ", emails)

	// defining and declaring a map
	fibonacciMap := map[int]int{
		1:  0,
		2:  1,
		3:  1,
		4:  2,
		5:  3,
		6:  5,
		7:  8,
		8:  13,
		9:  21,
		10: 34,
	}

	fmt.Println("Fibonacci Map: ", fibonacciMap)

}
