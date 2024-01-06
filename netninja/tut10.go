package main

import (
	"fmt"
	"strings"
)

// MULTIPLE RETURN VALUES

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)        // makes letter uppercase
	names := strings.Split(s, " ") // splits a string and in this case when there is a space ** " " **

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1]) // gets us 1st letter of each string
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_" // fn3,sn3

}

func main() {
	fn1, sn1 := getInitials("tifa lockhart")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("cloud strife")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("barret")
	fmt.Println(fn3, sn3)

}
