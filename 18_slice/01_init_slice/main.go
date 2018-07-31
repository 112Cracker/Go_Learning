package main

import (
	"fmt"
	)

/*
slice
	- definition
		- a data structure to store a list of values of a single type
	- reference
		- reference type -> a slice is referencing(points to) an underlying array
		- a pointer -> do not take & of a slice
		- length -> the number of elements the slice is currently holding,
		- capacity -> the length of the underlying array (capacity >= length)
	- the values of an uninitialized slice is nil
		- because it is a reference type
		- A initialized slice is always associated with an underlying array that holds its elements
		- slices are dynamic (unlike arrays)
			- length may change during execution
			for example, var my_slice []string{} -> not include capacity inside the bracket -> is a slice!!!
						 var my_array [42]string{} -> include capacity inside the bracket -> is a array!!!
		- The array underlying a slice may extend when slice length >= capacity
	- make
		- A slice created with make always allocates a new, hidden array to which the returned slice value refers
		make([]T, length, capacity)
		for example, make([]T, length, capacity) <=> new([capacity])[0:length]
		- slices may be composed to construct higher-dimensional slices
	- Attention to "index out of range errors!!!"
		- the length of a slice dictates the limit of accessible index
		can index a slice by 0 to len(my_slice) - 1
		for example, my_slice := make([]int, 5, 10)
					 my_slice[4] = 4 // valid assignment
					 my_slice[5] = 5 // index out of range error!!!
		instead, 	 my_slice = append(my_slice, 5)
	- deleting items from slices: use append !!!
		- my_slice = append(my_slice[:1], my_slice[2:]) // delete the second element from my_slice
	- incrementing a slice
		my_slice[0]++

	- creating a slice
		- shorthand
			- new_slice := []string{} (initialized)
		- var
			- sets slice to zero values which is nil
			var new_slice []string (uninitialized)
		- make
			new_slice := make([]string, length, capacity(optional))
			// if capacity not specified, capacity = length
*/

func main() {
	// shorthand slice initialization
	// with elements
	my_slice_not_empty := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%v is a %T\n", my_slice_not_empty, my_slice_not_empty) // []int slice type
	// without elements
	my_slice_empty := []int {}
	fmt.Printf("%v is a %T\n", my_slice_empty, my_slice_empty) // []int slice type
	// initialized but empty slice is not nil
	fmt.Println(my_slice_empty == nil) // -> false

	fmt.Println("------------------------------------------")

	my_array := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("%v is a %T\n", my_array, my_array) // [6]int array type

	// var slice initialization
	var my_slice_var []int
	fmt.Printf("%v is a %T\n", my_slice_var, my_slice_var)
	// uninitialized var slice is nil
	fmt.Println(my_slice_var == nil) // -> true

	fmt.Println("------------------------------------------")

	// make slice initialization
	my_slice_make := make([]int, 0, 2)
	fmt.Printf("%v is a %T that has length: %d and capacity: %d\n",
				my_slice_make, my_slice_make, len(my_slice_make), cap(my_slice_make))
	// slice is dynamic
	// capacity doubles each time length == capacity
	// capacity increase at some smaller rate after reaching some point
	for i:= 0; i< 42; i++ {
		my_slice_make = append(my_slice_make, i)
		fmt.Println("Len:", len(my_slice_make), "Capacity:", cap(my_slice_make), "Value: ", my_slice_make[i])
	}

}

