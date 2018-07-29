package main

import "fmt"

// named returns in Go
// specify the variable to return explicitly
// no need to append returned variable after the return statement
func greetings(first_name, last_name string) (greeting string) {
	greeting = fmt.Sprint("Hello ", first_name, " ", last_name, "!")
	return
}

func main() {
	fmt.Println(greetings("Harry", "Sun"))
}

/*
ATTENTION
	- Avoid using named returns
 */

/*
functions
purpose of functions
	- abstract code
	- code reusability
func receiver identifier(params) returns
parameters vs arguments
 */