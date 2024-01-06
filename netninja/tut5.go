package main

import "fmt"

func main() {

	// arrays *fixed length cannot be changed*
	// var ages [3]int = [3]int{20, 25, 30} // ****3 ITEMS ALL INT****
	var ages = [3]int{20, 25, 30}
	// ORDER: 0, 1, 2, 3 *STARTS AT 0*
	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60} // **NOT A FIXED LENGTH**
	scores[2] = 25
	scores = append(scores, 85) // OVERWRITES THE SCORES VAR

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3] // inclusive of the 1st number but not the 2nd number ex: [1:3]
	rangeTwo := names[2:]  // goes to end of array
	rangeThre := names[:3] // from start up to position 3
	fmt.Println(rangeOne, rangeTwo, rangeThre)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)
}
