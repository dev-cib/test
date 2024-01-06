package main

import "fmt"

func main() {
	y := 42
	z := 42.0
	fmt.Printf("%v is of type %T \n", y, y)
	fmt.Printf("%v is of type %T \n", z, z)

	var m float32 = 43.742
	fmt.Printf("%v is of type %T \n", m, m)

	/*
		// this does not work!
		// in Go you cant take a VALUE that is float32 and store it
		// in a var that is declared to hold a VALUE of float64
		z = m
		fmt.Printf("%v is of type %T \n", z, z)
	*/

	/*
		// this does work
		z = float64(m)
		fmt.Printf("%v is of type %T \n", z, z)
	*/
}
