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

	if x < 42 && y < 42 {
		// && requires both to be true to run
		fmt.Println("both are less than the meaning of life")
	}

	if x > 30 || x < 42 {
		// || requires one of them to be true to run
		fmt.Println("x is getting close to the meaning of life")
	}

	if x != 42 && y != 10 {
		fmt.Println("x is not 42 & y is not 10")
	}
	/*
		if x < 42 {
			fmt.Println("less than the meaning of life")
		} else {
			fmt.Println("equal to, or greater than, the meaning of life")
		}
	*/
}
