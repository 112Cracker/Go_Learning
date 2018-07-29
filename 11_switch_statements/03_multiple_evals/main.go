package main

import "fmt"

// Print today using multiple evaluations switch statements
func print_today(today string) {
	switch today {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Today is a weekday!")
	case "Saturday", "Sunday":
		fmt.Println("Today is weekend!")
	default:
		fmt.Println("Invalid Input")
	}
}

func main() {
	var today string
	for {
		fmt.Print("What day is today? ")
		fmt.Scan(&today)
		print_today(today)
	}

}
