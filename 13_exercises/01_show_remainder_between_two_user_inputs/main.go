package main

import "fmt"
/*
Write a program that allows the user to enter two numbers
then displays the remainder
when one number is divided by the other
 */
func main() {
	var a, b int
	fmt.Print("Please input the first number: ")
	fmt.Scan(&a)
	for b == 0 {
		fmt.Print("Please input the second number any integer but not a zero: ")
		fmt.Scan(&b)
	}
	fmt.Printf("The remainder of %d / %d is %d.", a, b, a%b)

}
