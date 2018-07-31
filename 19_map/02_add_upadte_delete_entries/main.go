package main

import "fmt"

func main() {
	myGrades := map[string]int {
		"Math": 95,
		"English": 96,
		"Physics": 92,
		"Chinese": 90,
	}
	fmt.Println(myGrades)

	// add entry
	myGrades["Chemistry"] = 89
	fmt.Println(myGrades)

	// update entry
	myGrades["Chinese"] -= 1
	fmt.Println(myGrades)

	// delete entry
	delete(myGrades, "Chemistry")
	fmt.Println(myGrades)

	fmt.Println("------------------------------------------")
	// update entry with exist checking
	if grade, exits := myGrades["History"]; exits {
		fmt.Printf("Wrong grade is %d: ", grade)
		myGrades["History"] += 1
	} else {
		myGrades["History"] = 89
	}
	fmt.Println(myGrades)

	fmt.Println("------------------------------------------")
	// delete no entry cause no error
	//delete(myGrades, "Music") // "Music" key not exists, but no error arises
	if _, exits := myGrades["Music"]; exits {
		fmt.Println("The student did not register for a Music class")
		delete(myGrades, "Music")
	} else {
		fmt.Println("Correct grades")
	}
}
