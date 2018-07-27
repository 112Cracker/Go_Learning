package main

import "fmt"

func wrapper() func() int {
	x := 0
	fmt.Println("RUN")
	return func() int {
		x++
		return x
	}
}

const p = "death & tax"

func main() {
	fmt.Println(p)
	const q = 42
	fmt.Println(q)


	fmt.Println(wrapper())
	increment := wrapper()
	fmt.Println(increment)

	fmt.Println(increment())
	fmt.Println(increment())
	//fmt.Println(increment())
}


