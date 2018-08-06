package main

import (
		"fmt"
)

// all types implement the empty interface
type pets interface {}

type dog struct {
	Id int
	Owner string
}

type cat struct {
	id int
	owner string
}

func main() {
	dog_0 := dog {
		0, "Jack",
	}

	dog_1 := dog {
		1, "Berry",
	}

	cat_0 := cat {
		1, "Sherry",
	}

	cat_1 := cat {
		2, "Mary",
	}

	pets := []pets{dog_0, dog_1, cat_0, cat_1}

	for index, pet := range pets {
		fmt.Println(index, pet)
	}

}
