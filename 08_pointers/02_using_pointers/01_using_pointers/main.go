package main

import "fmt"

func zero(x *int) {
	*x = 0
}

func main() {
	x := 42
	zero(&x)
	fmt.Println(x) // x is 0
}

/*
we can pass a memory address instead of a bunch of values (we can pass a reference)
and then we can change the value of whatever is stored at that memory address

we don't have to pass around large amounts of data
we only have to pass around addresses

Everything is PASSED BY VALUE in Go
When we pass a memory address, we are passing value too!!!
 */