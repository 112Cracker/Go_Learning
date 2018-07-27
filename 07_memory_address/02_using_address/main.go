package main

import "fmt"

const poundsToKg float64 = 0.45359

func main() {
	var pounds float64
	fmt.Print("Enter your weight in pounds: ")
	fmt.Scan(&pounds)
	kgs := pounds * poundsToKg
	fmt.Println("Your weight in kgs: ", kgs, "kg")
}