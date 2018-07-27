package main

import "fmt"

func main() {
	var a int
	var b string
	var c float64
	var d bool

	// %T get the type associated with the value
	// %v get the value associated with the initialization assignment
	fmt.Printf("%T \t %v \n", a, a)
	fmt.Printf("%T \t %v \n", b, b)
	fmt.Printf("%T \t %v \n", c, c)
	fmt.Printf("%T \t %v \n", d, d)
}

/*
var
	- var
		- zero value
 */

 /*
 Think about the difference between declaration, initialization and assignment
 var a int // is a declaration
 var a int = 42 // is a initialization
 a = 0 // is a assignment
  */
