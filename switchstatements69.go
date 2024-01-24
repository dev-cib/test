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

	switch {
	case x < 42:
		fmt.Println("x is LESS than 42")
	case x == 42:
		fmt.Println("x is EQUAL to 42")
	case x > 42:
		fmt.Println("x is GREATHER than 42")
	default:
		fmt.Println("this is the default case for ")
	}

	// ANOTHER WAY TO DO IT

	switch x {
	case 40:
		fmt.Println("x is 40")
	case 41:
		fmt.Println("X is 41")
	case 42:
		fmt.Println("X is 42")
	case 43:
		fmt.Println("x is 43")
	default:
		fmt.Println("this is the default case for x")
	}

	// fallthrough (GOES TO NEXT LINE)
	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough: x is 41")
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
	default:
		fmt.Println("printed because of fallthrough: this is the default case for x")
	}

	switch x {
	case 40:
		fmt.Println("x is 40")
		fallthrough
	case 41:
		fmt.Println("printed because of fallthrough statements: x is 41")
		fallthrough
	case 42:
		fmt.Println("printed because of fallthrough: x is 42")
		fallthrough
	case 43:
		fmt.Println("printed because of fallthrough: x is 43")
		fallthrough
	default:
		fmt.Println("printed because of fallthrough: this is the default case for x")
	}

}
