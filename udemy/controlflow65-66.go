package main

import "fmt"

func init() {
	fmt.Println("begin initialization")
}

func main() {
	// SEQUENCE
	fmt.Println("this is the first statement to run")
	fmt.Println("this is the second statement to run")
	x := 40 // this is the third statement to run
	y := 5  // this is the fourth statement to run
	fmt.Printf(" x=%v\n y=%v\n", x, y)

	// CONDITIONAL
	// if statements
	// switch statements

	if x < 42 {
		fmt.Println("less than the meaning of life")
	}

	if x < 42 {
		fmt.Println("less than the meaning of life")
	} else {
		fmt.Println("equal to, or greater than, the meaning of life")
	}
}
