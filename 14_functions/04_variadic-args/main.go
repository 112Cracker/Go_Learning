package main

import "fmt"

func averege(nums ...float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}


func main() {
	data := []float64{0, 1, 2, 3, 4, 5}
	// can have variadic arguments in Go
	// use ... to "open" an array(or slice)
	// and pass the argument to a func that have variadic parameters
	avg := averege(data...)
	fmt.Printf("The average is %0.2f", avg)
}