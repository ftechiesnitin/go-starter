package main

import "fmt"

func main() {
	// range for arrays
	ids := []int{1, 2, 3, 4, 5}

	// using index
	for i, id := range ids {
		fmt.Printf("%d - ID is: %d\n", i, id)
	}

	cities := []string{"paris", "london", "delhi", "amsterdam", "tokyo"}

	// without using index
	for _, city := range cities {
		fmt.Printf("City: %s\n", city)
	}

	// range for map
	capitals := map[string]string{"india": "delhi", "france": "paris", "singapore": "singapore", "germany": "berlin"}

	// key value pair
	for key, value := range capitals {
		fmt.Printf("%s: %s\n", key, value)
	}

	// only key
	for key := range capitals {
		fmt.Printf("Country Key: %s\n", key)
	}
	// fmt.Println("Hello World", capitals)
}
