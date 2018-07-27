package main

import "fmt"

/*
count the number of digits of a integer
and returns that number
 */
func countDigits(num *int) int {
	count := 0
	i := *num
	for i > 0 {
		i /= 10
		count ++
	}
	return count
}

/*
reverse the digits order of a integer
and return the revered number
non destructly
 */
func reverseDigits(num *int) int {
	i := *num
	reversed := 0
	for i > 0 {
		reversed = ((i % 10) + reversed) * 10
		i /= 10
	}
	reversed /= 10
	return reversed
}

func main() {
	var number int = 123
	fmt.Printf("The number of digits in %d is %d\n", number, countDigits(&number))
	fmt.Printf("The revered digits of %d is %d\n", number, reverseDigits(&number))
}
