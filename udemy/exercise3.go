package main

import "fmt"

/*
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "gophers")
	fmt.Println(a, b)

}
*/

func split(sum int) (x, y int) {
	x = sum * 4 / 9 // 17*4/9=7
	y = sum - x     // 17-7=10
	return x, y     // can do just "return" instead of "return x, y" which is a NAKED return
}

func main() {
	fmt.Println(split(17)) // 17 is the sum
}
