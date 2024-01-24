package main

import "fmt"

func main() {
	// SEQUENCE
	x := 40
	y := 5
	fmt.Printf(" x=%v \n y=%v \n", x, y)

	// LOOPS / INTERATIVE
	// for statements
	/*
		the go for loop is similar to-but not the same as-Cs
		it unifies for and while and there is no do-while
		there are 3 forms, only one of which has semicolons

		// Like a C for
		for init; condition; post { }

		// Like a C while
		for condition { }

		// Like a C for (;;)
		for { }
	*/
	// go.dev/doc/effective_go#for

	for i := 0; i < 5; i++ { // i++ adds 1 to i
		fmt.Printf("counting to five: %v \t first for loop \n", i)
	}

	for y < 10 {
		fmt.Printf("y is %v \t\t\t second for loop \n", y)
		y++
	}

	// break
	// takes you out of the loop
	for {
		fmt.Printf("y is %v \t\t third for loop \n", y)
		if y > 20 {
			break
		}
		y++
	}

	// continue
	// takes to next iteration
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue // prints 2 off and goes to 3 but thats not even so it continues to 4 which is even so that prints
		}
		fmt.Println("counting even numbers: ", i)
	}
}
