package main

import (
	"fmt"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("Nitin"))
	fmt.Println(getSum(4, 5))
}
