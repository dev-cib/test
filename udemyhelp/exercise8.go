package main

import "fmt"

// IOTA
const (
	c0 = iota // c0 == 0
	c1 = iota // c1 == 1
	c2 = iota // c2 == 2
)

/*
const (
	_ = iota // c0 == 0
	a = iota // c1 == 1
	b = iota // c2 == 2
	c = iota // c3 == 3
	d = iota // c4 == 4
	e = iota // c5 == 5
	f = iota // c6 == 6
)
*/

// ^ MINE - TODDS v

const (
	_ = iota
	a = iota
	b = iota
	c = iota
	d = iota
	e = iota
	f = iota
)

/*
const (

	c3 = iota // c0 == 0
	c4 = iota // c1 == 1
	c5 = iota // c2 == 2
	c6 = iota // c3 == 3

)
*/

func main() {

	fmt.Println(c0, c1, c2)
	fmt.Println(a, b, c, d, e, f)
	//fmt.Println(c3, c4, c5, c6)
	fmt.Printf("%d \t %b \n", 1, 1)
	fmt.Printf("%d \t %b \n", 1<<a, 1<<a)
	fmt.Printf("%d \t %b \n", 1<<b, 1<<b)
	fmt.Printf("%d \t %b \n", 1<<c, 1<<c)
	fmt.Printf("%d \t %b \n", 1<<d, 1<<d)
	fmt.Printf("%d \t %b \n", 1<<e, 1<<e)
	fmt.Printf("%d \t %b \n", 1<<f, 1<<f)
}
