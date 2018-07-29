package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var filtered []int
	for _, n := range numbers {
		if callback(n) {
			filtered = append(filtered, n)
		}
	}
	return filtered
}

func main() {
	evens := filter([]int{0, 1, 2, 3, 4}, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(evens) // [2 3 4]
}

// callback: passing a func as an argument