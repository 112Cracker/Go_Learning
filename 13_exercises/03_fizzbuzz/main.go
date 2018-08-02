package main

import "fmt"

/*
Write a program that prints the numbers from 1 to 100.
But for multiples of three print "Fizz" instead of the number
and for the multiples of five print "Buzz".
For numbers which are multiples of both three and five print "FizzBuzz".
 */
func fb_helper(number int) {
	switch {
	case number % 3 == 0 && number % 5 == 0:
		fmt.Printf("%d: FizzBuzz\n", number)
	case number % 3 == 0:
		fmt.Printf("%d: Fizz\n", number)
	case number % 5 == 0:
		fmt.Printf("%d: Buzz\n", number)
	}

}

func fizz_buzz() {
	upper := 100
	for i := 1; i <= upper; i++ {
		fb_helper(i)
	}
}

func main() {
	fizz_buzz()
}
