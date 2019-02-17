package main

import "fmt"

func main() {
	a := 5
	// b is a pointer to a (memory address)
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	// get the value from address
	fmt.Println(*b)
	// can also do
	fmt.Println(*&a)

	// change the value of variable using pointer
	*b = 10
	fmt.Println(a)
}
