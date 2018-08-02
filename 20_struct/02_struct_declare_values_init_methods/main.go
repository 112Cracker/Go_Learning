package main

import (
	"math"
	"fmt"
)

// complex struct
type complex struct {
	real float64
	img float64
}

// returns magnitude
func (c complex) calc_magitude() float64 {
	return math.Hypot(c.real, c.img)
}

// returns angle/phase
func (c complex) calc_phase() float64 {
	return math.Atan2(c.img, c.real)
}

// return a new complex whose value is c + b
func (c complex) plus(b complex) complex {
	sum := complex{}
	sum.real = c.real + b.real
	sum.img = c.img + b.img
	return sum
}

// return a new complex whose value is c - b
func (c complex) minus(b complex) complex {
	diff := complex{
		c.real - b.real,
		c.img - b.img,
	}
	return diff
}

// return a new complex whose value is alpha*c
func (c complex) scale(alpha float64) complex {
	scaled := complex{
		c.real * alpha,
		c.img * alpha,
	}
	return scaled
}

func main() {
	a := complex{1, 1}
	fmt.Printf("The magnitude of %0.0f*i+%0.0f is %0.2f\n", a.real, a.img, a.calc_magitude())
	b := complex{1, -1}
	fmt.Printf("The magnitude of %0.0f*i%0.0f is %0.2f\n", b.real, b.img, b.calc_magitude())

	c := a.plus(b)
	fmt.Printf("The sum of %v and %v is %v\n", a, b, c)

	d := a.minus(b)
	fmt.Printf("The difference of %v and %v is %v\n", a, b, d)

}