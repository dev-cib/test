package main

import "fmt"

func main() {

	age := 23
	name := "christian"

	// Print *doesnt add new line whereas Println does*
	fmt.Print("test, ")
	fmt.Print("score? \n") // \n *new line*
	fmt.Print("100 \n")

	// Println
	fmt.Println("new test")
	fmt.Println("new score?")
	fmt.Println("100")

	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings) *** %_ = FORMAT SPECIFIER *** %v is the default format for variables
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name) // %q wont work on int *ideally used on strings* // %q double "s around var
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %f points on your test! \n", 225.55)    // %f for float
	fmt.Printf("you scored %0.1f points on your test! \n", 225.55) // limit decimal points by putting a number between % f // ex: 1 decimal = 0.1 / 2 decimal = 0.2
	fmt.Printf("you scored %0.2f points on your test! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}
