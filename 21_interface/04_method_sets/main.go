package main

//
import (
	"fmt"
	"math"
)

/*
Receivers       Values
-----------------------------------------------
(t T)           T and *T
(t *T)          *T

Values          Receivers
-----------------------------------------------
T               (t T)
*T              (t T) and (t *T)
 */
// Use value receiver to avoid mistakes?

type circle struct {
	radius float64
}

type shape interface {
	area() float64
	dimension() string
}

// value receiver
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// pointer receiver
func (c *circle) dimension() string {
	return "2D"
}

func info(s shape) {
	fmt.Println("area", s.area())
	fmt.Println("dimension", s.dimension())
}

func main() {
	c := circle{5}
	// value receiver
	//info(c) // value
	info(&c) // pointer
}