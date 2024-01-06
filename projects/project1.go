package main

import "fmt"

func main() {

	gender := "boy"
	time := 4
	dad := "christian"
	mom := "mikayla"
	// age := 23
	ages := []int{23, 24}
	// age = 24 // ask dad for help
	amber := "mimi"
	eric := "papaw"

	var baby = fmt.Sprintf("%v and %v are the mother and father. they are ages %v", mom, dad, ages)
	fmt.Println("the parents:", baby)
	fmt.Printf("the baby %v will be arriving in %v months \n", gender, time)
	fmt.Printf("his %v is amber and his %v is eric. these are the parents of %v \n", amber, eric, dad)

}
