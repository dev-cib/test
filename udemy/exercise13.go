package main

import "fmt"

func main() {
	// reminder: uint = unsigned int
	// int = signed int

	var a int8 = 127
	fmt.Printf("%v is of type %T \n", a, a)
	var b uint8 = 255
	fmt.Printf("%v is of type %T \n", b, b)

}

/*
write a program that delcares two variables
	- one variable to store a VALUE of TYPE int8 (assign largest number and print)
and
	- one variable to store a VALUE of TYPE uint8 (assign largest number and print)
*/
