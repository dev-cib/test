package main

import (
	"fmt"
	"math/rand"
)

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

	/*
		the expression [evaluated in an if statement] may be preceded by a simple statement
		which executes before the expression is evaluated
	*/
	// go.dev/ref/spec#If_statements
	/*
		the comma ok idiom is also like this
	*/
	// go.dev/play/p/OXGzjxVkag0

	// STATEMENT;STATMENT IDIOM
	if z := 2 * rand.Intn(x); z >= x { // rand.Intn(random integer) and since x = 40 itll be a random # 0-40 multiplied by 2 = value of z
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}
}
