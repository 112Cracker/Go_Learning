package main

import (
	"fmt"
	"strconv"
)

func main() {
	// int to float

	// float to int

	// rune to string
	var x rune = 'a' // rune is an alias for int32; normally omitted in this statement
	fmt.Println(x)
	fmt.Println(string(x))
	// conversion: rune to string


	// rune to []bytes
	fmt.Println([]byte{'h', 'e', 'l', 'l', 'o'}) //[104 101 108 108 111]

	// string to []byte
	fmt.Println([]byte("hello")) //[104 101 108 108 111]

	// rune to []bytes to string
	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'})) //hello

	// strconvƒ
	var num1 = "42"
	var num2 = 42
	num3, _ := strconv.Atoi(num1)
	fmt.Println(num2==num3) // true

	num4 := strconv.Itoa(num2)
	fmt.Println(num1==num4) // true

	// ParseBool, ParseFloat, ParseInt, and ParseUnit
	a, _ := strconv.ParseBool("true")
	b, _ := strconv.ParseFloat("3.14", 64)
	c, _ := strconv.ParseInt("-42", 10, 64)
	d, _ := strconv.ParseUint("42", 10, 64)

	fmt.Println(a, b, c, d)

	// FormatBool, FormatFloat, FormatInt, and FormatUint
	e := strconv.FormatBool(true)
	f := strconv.FormatFloat(3.14, 'E', -1, 64)
	g := strconv.FormatInt(42, 2)
	h := strconv.FormatUint(42, 16)

	fmt.Println(e, f, g, h)
}

