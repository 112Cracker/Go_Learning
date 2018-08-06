package main

import (
	"fmt"
)

type dog struct {
	Id int
	Owner string
	IsFriendly bool
}

type cat struct {
	Id int
	Owner string
	IsFriendly bool
}

// a interface{} -> accept any type parameter
func info(a interface{}) string {
	return fmt.Sprint(a)
}


func main() {
	dog_0 := dog {
		0, "Jack", true,
	}

	cat_0 := cat {
		1, "Sherry", false,
	}

	var test_0 int
	var test_1 float64

	fmt.Println(info(dog_0))
	fmt.Println(info(cat_0))
	fmt.Println(info(test_0))
	fmt.Println(info(test_1))

}
