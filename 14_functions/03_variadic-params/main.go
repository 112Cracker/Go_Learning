package main


import "fmt"

// can have variadic parameters in Go
// various number of parameters
// compare to *args in Python
func averege(nums ...float64) float64 {
	var sum float64
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}


func main() {
	avg := averege(0, 1, 2, 3, 4, 5)
	fmt.Printf("The average is %0.2f", avg)
}