package main

import (
	"fmt"
	"math"

	"github.com/ftechiesnitin/go-starter/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.8))
	fmt.Println(math.Ceil(2.5))
	fmt.Println(math.Sqrt(4))
	fmt.Println(strutil.Reverse("I am awesome at go development"))
}
