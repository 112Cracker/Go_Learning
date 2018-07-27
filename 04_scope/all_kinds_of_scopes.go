package main

import "fmt"

// wrapper is a func that returns a func!!!
func wrapper() func() int {
	x := 0 // x is visible to the returned func because it is in the outer space
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper() // assign the returned func from wrapper to a variable called increment
	fmt.Println(increment()) // => 1
	fmt.Println(increment()) // => 2

	fmt.Println(y)
	fmt.Println(password)
}

var y = 42 // y is visible to all the func(s) inside package

/*
scope
	- levels of scope
		- universe
		- package
		- file
		- block (things within a curly brace {})
	- keep scope tight
	- closure
		- variables declared (initialized) in the outer scope are visible to the statements
		  in the inner scopes which are enclosed by the outer scope
 */

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
