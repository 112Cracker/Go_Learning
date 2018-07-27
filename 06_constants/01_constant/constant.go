package main

import "fmt"

// can have single declaration
//const name = "Harry"

// can have multiple declarations
const (
	pi = 3.14
	e = 2.71
)

func main() {
	fmt.Printf("pi is %f\n", pi)
	fmt.Printf("natural log is %f\n", e)
}

/*
Constants
	- a simple, unchanging value
	- a parallel type system
		- there are TYPED and UNTYPED constants
			- const greeting = "Hello"
			- const typedGreeting string = "Hello"
	- UNTYPED constant
		- a constant value that does not have a fixed type
			- a "kind"
			- not yet forced to obey strict rules that prevent combining differently typed typed values
 */