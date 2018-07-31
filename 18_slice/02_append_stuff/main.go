package main

import "fmt"

func main() {
	greetings := make([]string, 1, 2)
	// 1 is the length -> number of elements pointed to by the slice
	// 2 is the capacity -> number of elements in the underlying array

	greetings[0] = "Good morning!"
	//greetings[1] = "Bonjour!" //panic: runtime error: index out of range
	// append to access and assign beyond slice scope(length) limit
	greetings = append(greetings, "Bonjour!")
	greetings = append(greetings, " Buenos dias! ")
	fmt.Printf("%v has length: %d and capacity: %d\n",
		       greetings, len(greetings), cap(greetings))

	// ATTENTION: Be aware of the slice length
	// Try to avoid index beyond the slice length
	//third_greeting := greetings[3] //panic: runtime error: index out of range
	//fmt.Println(third_greeting)

	fmt.Println("----------------------------------------------")

	first_slice := []int{0, 1, 2}
	second_slice := []int{3, 4, 5}

	// use ... to "open" a slice
	merged_slice := append(first_slice, second_slice...)
	// the same as
	merged_slice_another := append(first_slice, 3, 4, 5)

	fmt.Printf("%v and %v are merged into %v\n",
				first_slice, second_slice, merged_slice)

	fmt.Printf("%v and %v are merged into %v\n",
		first_slice, second_slice, merged_slice_another)

	fmt.Println("----------------------------------------------")
	// use append to delete elements from a slice
	first_slice = append(first_slice[0:1], first_slice[2:]...)
	fmt.Print(first_slice) // the second element is deleted
}

