package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like my hat?`
	g := 'M'

	// %T get the type associated with the value
	// %v get the value associated with the initialization assignment
	fmt.Printf("%T \t %v \n", a, a)
	fmt.Printf("%T \t %v \n", b, b)
	fmt.Printf("%T \t %v \n", c, c)
	fmt.Printf("%T \t %v \n", d, d)
	fmt.Printf("%T \t %v \n", e, e)
	fmt.Printf("%T \t %v \n", f, f)
	fmt.Printf("%T \t %v \n", g, g)


}
/*
var
	- shorthand
		- can only be used inside a func
 */