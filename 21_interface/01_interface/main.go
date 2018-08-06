//CITE CODE: FROM https://github.com/GoesToEleven/GolangTraining/blob/master/21_interfaces/01_interface/04_interface/main.go
package main

import (
	"math"
	"fmt"
)

// define a interface named shape
// area() method will calculate and return the area of the given shape
type shape interface {
	area() float64
}


// define a struct named square
type square struct {
	side float64
}


func (s square) area() float64 {
	return s.side * s.side
}

// define a struct names circle
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}


// method that takes a shape
func calc_shape_area(s shape) float64 {
	return s.area();
}

func main() {
	s := square{2}
	fmt.Printf("This square's area is %0.2f\n", calc_shape_area(s))

	c := circle{2}
	fmt.Printf("This circle's area is %0.2f\n", calc_shape_area(c))
}
