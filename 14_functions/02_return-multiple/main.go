package main

import "fmt"

const mile_to_km = 1.60934

// can have multiple returns in Go

// Calculate distance using speed and hours
// And return the distance in both miles and km
func calc_distance(speed float64, hours float64) (float64, float64){
	miles := speed * hours
	kms := miles * mile_to_km
	return miles, kms
}

func main() {
	speed := 70.0
	hours := 2.5
	// can have parallel assignments in Go
	miles, kms := calc_distance(speed, hours)
	fmt.Printf("Distance travled is %0.2f miles or %0.2f km", miles, kms)
}

