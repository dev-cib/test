package main

import "fmt"

func main() {
	christian := 24
	fmt.Printf("24 as binary, %b \n", christian)
	fmt.Printf("24 as hexadecimal, %x \n", christian)
	// %x hexideciaml notation
	// %b binary notation

	// MY WAY
	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5
	fmt.Printf("binaries of a-f: %b, %b, %b, %b, %b, %b \n", a, b, c, d, e, f)
	fmt.Printf("hexadeciamls of a-f: %x, %x, %x, %x, %x, %X \n", a, b, c, d, e, f)

	// TODD
	fmt.Printf("%v \t %b \t %x \n", a, a, a)
	fmt.Printf("%v \t %b \t %x \n", b, b, b)
	fmt.Printf("%v \t %b \t %x \n", c, c, c)
	fmt.Printf("%v \t %b \t %x \n", d, d, d)
	fmt.Printf("%v \t %b \t %x \n", e, e, e)
	fmt.Printf("%v \t %b \t %x \n", f, f, f)

	//exercise from todd: print these values as both binary and hexadecimal
	// a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

}
