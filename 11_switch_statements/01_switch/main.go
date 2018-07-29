package main

import "fmt"


func print_todo_on(today string) {
	switch today {
	case "Monday":
		fmt.Println("Fresh up, start of a new week!")
	case "Tuesday":
		fmt.Println("Meeting scheduled for today")
	case "Wednesday":
		fmt.Println("Do not forget to watch the football game today!")
	case "Thursday":
		fmt.Println("An ordinary weekday")
	case "Friday":
		fmt.Println("Have a good weekend!")
	case "Saturday":
		fmt.Println("Enjoy weekend!")
	case "Sunday":
		fmt.Println("Enjoy weekend! ")
	}
}

func main() {
	var today string
	for {
		fmt.Print("What day is today? ")
		fmt.Scan(&today)
		print_todo_on(today)
	}

}

/*
no default fallthrough
can specify a fallthrough by explicitly stating a fallthrough
do not need a break
 */