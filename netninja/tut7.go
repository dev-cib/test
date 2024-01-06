package main

import "fmt"

func main() {

	// LOOPS
	x := 0
	for x < 5 {
		fmt.Println("the value of x is:", x)
		x++ // adds 1 to x
	}
	// BOTH THE SAME
	for i := 0; i < 5; i++ {
		fmt.Println("the value of i is:", i)

	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for a := 0; a < len(names); a++ {
		fmt.Println(names[a])
	}

	//for index, value := range names {
	//	fmt.Printf("the value at index %v is %v \n", index, value)
	//}

	for _, value := range names { // _ allows me to not use the var without erroring
		fmt.Printf("the value is %v \n", value)
		value = "new string" // doesnt do anything **************
	}

	fmt.Println(names)

}
