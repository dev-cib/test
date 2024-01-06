package main

import "fmt"

const yoshi = "yoshi"

func main() {

	// strings - strings are double quotes *number = integer*
	var nametw0 string = "luigi"
	var namethr33 = "mario"
	var namef0ur string

	fmt.Println(nametw0, namethr33, namef0ur)

	nametw0 = "toad"
	namef0ur = "bowser"

	fmt.Println(nametw0, namethr33, namef0ur)
	// cant use shorthand oustide of func ex: :=
	// := declares a var and assigns its value in a single statement
	namefive := "yoshi"
	// namefive = christian
	fmt.Println(namefive)
	fmt.Println(yoshi)

	// integers (ints) is whole # // floats is decimals
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	//int8 range: -128 to 127
	var numOne int8 = 127
	var numTwo int8 = -128
	var numThree uint = 25  //uint cant have a negative number
	var numFour uint8 = 255 //uint8 range: 0 to 255 *NO NEGATIVES*
	// **FLOAT IS DECIMAL**
	var scoreOne float32 = -25.98               //uses range of int32: -214748.3648 to 214748.3647
	var scoreTwo float64 = 883247729784397812.9 // ^ int64
	scoreThree := 1.5

}
