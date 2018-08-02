package main

import "fmt"

/*
write a program that loops from 1 - range
printing even numbers
 */

func print_evens_lq_than(upper int) {
	for i := 1; i <= upper; i++ {
		if remainder := i%2; remainder == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	input := -1
	for input < 0 {
		fmt.Print("Find even integers less than: ")
		fmt.Scan(&input)
		if (input < 0) {
			fmt.Println("Please input a positive integer!")
		}
	}
	print_evens_lq_than(input)
}

