package main

import (
	"fmt"
	)

// init; condition; post -> for loop in Go
// Print positive integers less than num using ordinary for loop
func printLessThan_forLoop(num int) {
	for i:= 0; i < num; i++ {
		fmt.Println(i)
	}
}

// nested loops
// Print 9 by 9 multiplication charts in the console
func printMultCharts() {
	for i:= 1; i < 10; i++ {
		for j:= 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d\t", j, i, j*i)
		}
		fmt.Println();
	}
}

// condition -> while loop in Go
// Print positive integers less than num using while loop
func printLessThan_whileLoop(num int) {
	i := 0
	for i < num {
		fmt.Println(i)
		i ++
	}
}

// including break or continue
// for with no condition -> infinite loop in Go
func printLessThan_includeBreak(num int) {
	i := 0
	// for with no condition is infinite loops in Go
	for {
		if i >= num {
			break
		}
		fmt.Println(i)
	}
}


func main() {
	printLessThan_forLoop(100)
	printLessThan_whileLoop(100)
	printLessThan_includeBreak(100)

	printMultCharts()
}

/*
Control Flow
	- sequence
	- loop/iterative
		- nested loops
		- break
		- continue
 */