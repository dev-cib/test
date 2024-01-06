package main

import "fmt"

/*
var c, python, java bool
const d int = 42
*/

// var i, j int = 1, 2
var i, j = 1, 2.0

func main() {
	//	var i
	//	fmt.Println(i, c, python, java, d)

	var c, python, java = true, false, "no!"
	// fmt.Println(i, j, c, python, java)
	fmt.Printf("%T \t %T \t %T \t %T \t %T \t", i, j, c, python, java)
}
