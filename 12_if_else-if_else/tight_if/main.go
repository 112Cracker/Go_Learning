package main

import "fmt"

func is_even(num int) bool{
	if remainder := num % 2; remainder == 0 {
		return true
	}
	return false
}


func main() {
	if is_even(0) {
		fmt.Println("It is a even number")
	} else {
		fmt.Println("It is a odd number")}
}

/*
Could make a tight scope
Inline variable only accessible within the scope
 */

