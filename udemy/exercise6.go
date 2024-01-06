package main

import "fmt"

func main() {
	/*
		i := 42 // int
		f := 3.142 // float64
		g := 0.867 + 0.5i // complex128
	*/

	var i int
	j := i                                  // j is an int
	fmt.Printf("%v is of type %T \n", j, j) // j is 0 because its a zero value

	// v := 42 // change me!
	v := "42"
	fmt.Printf("%v is of type %T \n", v, v)

}
