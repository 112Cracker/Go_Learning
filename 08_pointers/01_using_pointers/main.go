package main

import "fmt"

func main() {

	x := 42

	fmt.Println(x) // => 42
	fmt.Println(&x) //

	var y = &x
	fmt.Println(y) // => 42
	fmt.Println(&y)

	*y = 0
	fmt.Println(x) // => 0
}

/*
Pointers
	- Using memory address in statements
	- referencing / dereferencing
	- It's all passed by value

 */