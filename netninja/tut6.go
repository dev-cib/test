package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))         // returns "true" since the greeting *contains* hello
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) // replaces hello with hi
	fmt.Println(strings.ToUpper(greeting))                   // uppercase
	fmt.Println(strings.Index(greeting, "ll"))               // ex: hello - starts at 0. h0,e1,l2,l
	fmt.Println(strings.Split(greeting, " "))                // array
	// *** THE ORIGINAL VALUE IS UNCHANGED ***
	fmt.Println("original string value =", greeting)

	ages := []int{45, 20, 21, 22, 23, 24, 50, 60}

	sort.Ints(ages) // does what it says: sorts ages in numerical order
	fmt.Println(ages)

	index := sort.SearchInts(ages, 23) // **this searches the SORTED ages**
	fmt.Println(index)

	index2 := sort.SearchInts(ages, 90) // 90 isnt in the slice but it still shows where it would be
	fmt.Println(index2)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}
	sort.Strings(names) // sorts name in alphabetical order
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "bowser")) // searches the SORTED names

}
