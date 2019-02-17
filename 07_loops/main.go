package main

import "fmt"

func main() {

	// for loop long method
	i := 1

	for i <= 3 {
		fmt.Println("i : ", i)
		i++
	}

	// for loop short method
	for j := 0; j < 3; j++ {
		fmt.Println("j : ", j)
	}

	// Fiz buzz challenge
	for k := 1; k <= 100; k++ {
		if k%15 == 0 {
			fmt.Println("FizzBuzz", k)
		} else if k%3 == 0 {
			fmt.Println("Fizz", k)
		} else if k%5 == 0 {
			fmt.Println("Buzz", k)
		}
	}
	fmt.Println("Hello World")
}
