package main

import "fmt"

var a int = 24

func main() {
	me := "Christian"
	fmt.Printf("%v is %v \n", me, a)

	var b float64 = 24.24
	fmt.Println(b)

	fmt.Printf("%v is of type %T \n", a, a)
	fmt.Printf("%v is of type %T \n", me, me)
	fmt.Printf("%v is of type %T \n", b, b)

	// ^ MINE - TODDS v

	s, i, f := "happiness", 42, 42.42
	fmt.Printf("%v is of type %T \n", s, s)
	fmt.Printf("%v is of type %T \n", i, i)
	fmt.Printf("%v is of type %T \n", f, f)

}

/*
var storing a value of type
 - string
 - int
 - float64
use print verbs to show
 - value
 - type
*/
