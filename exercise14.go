package main

import "fmt"

var baby string = "Theodore"

const (
	age  = 32
	time = "weeks"
)

func main() {

	dad := "Christian"
	mom := "Mikayla"

	fmt.Printf("%v is %v %v old. %v's dad and mom are %v and %v \n", baby, age, time, baby, dad, mom)

}

/*
create the following vars with the following scopes:
 
- package level
	(create outside of func main)
	(use the)
		var keyword
		const keyword
 	
	
 - block level
	(inside the func main)
	(use short declaration operator)

USE THE VARIABLES IN FUNC MAIN
/*