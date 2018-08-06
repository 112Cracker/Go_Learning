package main

import (
	"fmt"
	"reflect"
)

// user-defined struct type comparison in go
type person struct {
	name string
	id int
}

// use interface casting to implement deepEqual
// for user-defined struct type
func (p person) equal(other interface{}) bool {
	_, isPerson := other.(person)
	if isPerson {
		// if it is a person
		// cast to person before comparison fields
		return other.(person).name == p.name &&
			   other.(person).id == p.id
	}
	return false
}

func main() {
	tom := person {
		"Tom",
		101,
	}

	jerry := person {
	"Jerry",
	102,
	}

	tommy := person {
		"Tom",
		101,
	}

	var jenny interface{}

	// tom and jerry are different person
	fmt.Println(reflect.DeepEqual(tom, jerry)) // => false
	fmt.Println(tom.equal(jerry)) // => false
	fmt.Println(tom.equal(jenny)) // => false
	// tom and tommy are the same person
	fmt.Println(reflect.DeepEqual(tom, tommy)) // => true
	fmt.Println(tom.equal(tommy)) // => true

	var unknown interface{} = "mystery"
	str, isString := unknown.(string)
	if isString {
		fmt.Printf("%T: %v\n", str, str)
	} else {
		fmt.Println("Invalid input")
	}
}
