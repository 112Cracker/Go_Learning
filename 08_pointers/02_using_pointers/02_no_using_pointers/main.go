package main

import "fmt"

func zero(x int) {
	x = 0
}

func main() {
	x := 42
	zero(x)
	fmt.Println(x) // x is still 42
}