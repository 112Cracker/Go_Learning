package main


import (
	"fmt"
	"github.com/112Cracker/GoLearning/02_package/stringutil"
)

func main() {
	fmt.Println(stringutil.UserName)
	//fmt.Println(stringutil.username)

	num := 13
	if stringutil.IsPrime(num) {
		fmt.Printf("%d is a prime numer", num)
	}

	//if stringutil.isEven(num) {
	//	fmt.Printf("%d is a even number", num)
	//}
}

/*
Packages
- one directory containing files
	- package declaration in every file
	- package scope
		- something in one file are accessible to others
- exported / unexported
	- in other words, visible or invisible
	- generally, do not user public or private on Golang
	- capitalization
		- capitalized variables, functions etc. are visible outside the package scope -> exported
		- lowercase ... not visible outside the package scope -> unexported
*/