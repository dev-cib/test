package main

// MORE COMPLEX (EXAMPLE)

import "fmt"

var Celsius float64
var Farenheit float64

func main() {
	fmt.Print("enter temperature(F): ")
	// fmt.Scanf("%f", &Farenheit) // ask dad to explain // why does %b work but %d doesnt
	fmt.Scan(&Farenheit) // morts example
	Celsius = (Farenheit - 32) * 5 / 9
	fmt.Printf("temperature F in C: %0.2f \n", Celsius)
}

/*
declared var Farenheit and Celsius to be of type float64
prints "enter temperature(F)" to user in terminal
user puts in a number(F)
?????
clesius is assigned the formula for F to C
prints back to the user "temperature F in C: %v"
%v = the temp in C (varies depending on F input)
*/
