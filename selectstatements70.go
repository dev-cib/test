package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ** PREVIEW **

func main() {
	// CONDITIONAL
	// if statements
	// switch statements

	// concurrency
	// select a channel

	ch1 := make(chan int)
	ch2 := make(chan int)

	// get an int64, convert it to the type time.Duration, then use it in time.Sleep
	// Int63n returns an int64
	// type durations underlying type as int64
	// time.Sleep takes type duration
	// time.Millisecond is a constant in the time package
	// pkg.go.dev/time#pk-constants
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))
	// fmt.Printf("%v \t %T \n", d1, d1)
	// time.Sleep(d1 * time.Millisecond)
	// fmt.Println("Done")

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- 41
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	// a "select" statement chooses which of a set of possible send or recieve operations will proceed
	// it looks similar to a "switch" statement but with the cases al referring to comunication operation
	// go.dev/ref/spec#Select_statements
	select {
	case v1 := <-ch1:
		fmt.Println("value from channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}

}
