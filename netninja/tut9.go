package main

import (
	"fmt"
	"math"
)

// USING FUNCTIONS

func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}
func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// r = radius
func circleArea(r float64) float64 {
	return math.Pi * r * r // formula for AREA OF A CIRCLE. A = pir^2
	// return the area of circle
}

func main() {

	//sayGreeting("mario")
	//sayGreeting("luigi")
	//sayBye("mario")

	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting)
	cycleNames([]string{"cloud", "tifa", "barret"}, sayBye)

	a1 := circleArea(10.5) // ** RETURNS VALUE r STORES IT IN a1 **
	a2 := circleArea(15)

	fmt.Println(a1, a2) // area of the 2 circles a1/a2
	fmt.Printf("the area of the 2 circles are: %0.1f and %0.1f \n", a1, a2)

}
