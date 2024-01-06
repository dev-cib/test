package main

// MINE

import "fmt"

var Farenheit int

func main() {
	Farenheit = 24
	Celsius := (Farenheit - 32) * 5 / 9
	fmt.Printf("%vF is %v in Celsius \n", Farenheit, Celsius)

}

/*
declared var Farenheit to be of type int
assigned 24 to var Farenheit
Celsius is assigned the forumla for F to C
var Farenheit (24) is plugged into the equation
Farenheit is assigned in the 1st %v and clesius is assigned the 2nd %v
*/
