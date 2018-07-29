package main

import "fmt"

// can have slice as parameters
func averege(nums []float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}


func main() {
	data := []float64{0, 1, 2, 3, 4, 5}
	avg := averege(data)
	fmt.Printf("The average is %0.2f", avg)
}

/*
Ways to pass in multiple arguments(same type) as the same time
	- have variadic parameters function
		- and pass in variadic arguments
		- or pass in slices as variadic arguments (data...)
	- have slice parameter function
		- and pass in slices
 */