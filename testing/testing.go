package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 100

	if y := 4 * rand.Intn(x); y < x {
		fmt.Printf("y is %v and that is less than x which is %v\n", y, x)
	} else {
		fmt.Printf("y is %v and that is greater than x which is %v\n", y, x)
	}
}
