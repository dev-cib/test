package main

import "fmt"

// HINT @ EFFECTIVE GO "IOTA"
type Bytesize int

/*
const (
	_           = iota // throwing away the 0
	KB Bytesize = 1 << (10 * iota)
	MB Bytesize = 2 << (10 * iota)
	GB Bytesize = 3 << (10 * iota)
	TB Bytesize = 4 << (10 * iota)
	PB Bytesize = 5 << (10 * iota)
	EB Bytesize = 6 << (10 * iota)
)
*/

// ^ MINE - TODDS v

const (
	_           = iota
	KB Bytesize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Printf("%d \t %b \n", KB, KB)
	fmt.Printf("%d \t %b \n", MB, MB)
	fmt.Printf("%d \t %b \n", GB, GB)
	fmt.Printf("%d \t %b \n", TB, TB)
	fmt.Printf("%d \t %b \n", PB, PB)
	fmt.Printf("%d \t %b \n", EB, EB)

}
