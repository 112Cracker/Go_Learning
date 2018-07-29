package main

import "fmt"

/*
Print printable characters from ASCII
 */
func main() {
	for i := 32; i <= 126; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
}

/*
terminology
	- lexical elements
	- literal values
	- runes

rune
	- character
	- an integer value identifying a Unicode code point
	- alias for int32
		- UTF-8 is a 4 byte coding schema
		- with int32(4 bytes) we have numbers for each of the code points
 */