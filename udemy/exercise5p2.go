package main

import (
	"fmt"
	"math"
)

func main() {
	/* var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	*/

	// better way to write (i did this)
	x := 3
	y := 4
	f := float64(math.Sqrt(float64(x*x + y*y)))
	z := uint(f)
	fmt.Println(x, y, z)

	/*
		NUMERIC CONVERSIONS:
		var i int = 42
		var f float64 = float64(i)
		var u uint = uint(f)

		*BETTER WAY TO WRITE THE ABOVE* (simply)
		i := 42
		f := float64(i)
		u := uint(f)
	*/
}
