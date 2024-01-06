package main

import "fmt"

const Pi = 3.14 // CONST cannot be declared using := (short declaration) syntax

func main() {
	const Hello = "こんにちは"
	fmt.Println(Hello, "World")
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

}
