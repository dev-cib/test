package main

import "fmt"

func main() {
	// BOOLEANS & CONDITIONALS
	age := 45
	fmt.Println(age <= 50) // age is less than or equal to 50
	fmt.Println(age >= 50) // age is greater than or equal to 50
	fmt.Println(age == 45) // age is equal to 45
	fmt.Println(age != 50) // ! = not equal to 50

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 { // moves to else clause if above expression isnt true
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue // ****SAYS GO BACK UP TO THE FOR LOOP AT THE START DONT CONTINUE DOWN BELOW (printf) but continue with the LOOP****
		}

		if index > 2 {
			fmt.Println("breaking at pos", index)
			break // ****SAYS BREAK OUT OF THE LOOP COMPLETELY SO DONT EVEN GO BACK TO THE TOP AND DONT CONTINUE CYCLING THROUGH REST OF THE SLICE****
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}

}
