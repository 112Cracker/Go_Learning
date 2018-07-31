package main

import "fmt"

func main() {
	// iterate the map using key value pair
	// use range on map
	myGrades := map[string]int {
		"Math": 95,
		"English": 96,
		"Chinese": 90,
	}
	fmt.Printf("Subject\tGrade\n")
	for subject, grade := range myGrades {
		fmt.Printf("%s: %d\n", subject, grade)
	}
}
