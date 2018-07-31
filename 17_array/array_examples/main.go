//CITE CODE: FROM https://github.com/GoesToEleven/GolangTraining/blob/master/17_array/05/main.go

package main

import "fmt"

func main() {
	// byte alias for unit8
	var x [256]byte

	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 0; i < 256; i++ {
		// byte(i) conversion from int to byte type
		x[i] = i
	}

	for i, v := range x {
		// %v -> value; %T -> type; %b -> binary
		fmt.Printf(": %v - %T - %b\n", v, v, v)
	}
}

/*
array
	- initialization
		- var identifier [size]type
	- len(array) returns the size of the array
	- use range to iterate the array type
		- for index, value := range array {
						...
			}
 */

