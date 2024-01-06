package main

import "fmt"

func main() {
	// defined some constants with identifier of name and age and assigned values to those constants. ex: name "christian", age 23
	// could write it like lines 8-9 or line 11
	// const name = "christian"
	// const age = 23

	const name, age = "christian", 23

	// fmt.Printf("%v is %d years old. \n", name, age) // could use %s for string instead of %v variable
	fmt.Printf("%v is %d years old. \t and the type is %T and %T \n", name, age, name, age)
}
