package main

import "fmt"

func main() {

	a := 42 // := SHORT DECLERATION OPERATOR: allows the user to quickly assign a value to a variable
	fmt.Println(a)

	var g int
	fmt.Println(g) // returns a 0 by default since g is of type int but no value is given

	g = 43
	fmt.Println(g)

	b, c, d, _, f := 0, 1, 2, 3, "happiness" // doesnt print 3 because its assigned to the blank identifier _
	fmt.Println(b, c, d, f)

	// this would not work
	/*
		b, c, d, e := 0, 1, 2, 3
		fmt.Println(b, c, d) // e is not being used (would need to use blank identifier _ instead of e)
	*/

	// VAR + example
	/*
		// DECLARE a variable to hold a VALUE of a certain TYPE
		// then ASSIGN a VALUE of that TYPE to the variable
		// initializing a variable
		var h int = 44
		fmt.Println(h)
	*/

}
