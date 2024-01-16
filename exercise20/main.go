package main

import (
	"fmt"

	"github.com/dev-cib/puppy"
)

func main() {
	fmt.Println("hello gophers")
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.BigBark()
	s4 := puppy.BigBarks()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	puppy.From12()
}

// wanted me to downgrade to puppy@v1.2.0 and run
